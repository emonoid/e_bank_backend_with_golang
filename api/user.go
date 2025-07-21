package api

import (
	"database/sql"
	"log"
	"net/http"

	db "github.com/emonoid/islami_bank_go_backend/db/sqlc"
	"github.com/emonoid/islami_bank_go_backend/utils"
	"github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	UserName string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	UserName string `json:"username"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		UserName: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.UserName,
		HashedPassword: hashedPass,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, gin.H{"error": "this user already exist"})
			}
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	userResponse := newUserResponse(user)

	ctx.JSON(http.StatusOK, userResponse)
}

type getUserRequest struct {
	UserName string `uri:"username" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)

}

type logineUserRequest struct {
	UserName string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req logineUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = utils.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)

}

// type updateAccountRequest struct {
// 	ID        int64  `json:"id" binding:"required"`
// 	OwnerName string `json:"owner" binding:"required"`
// 	Currency  string `json:"currency" binding:"required,currency"`
// 	Balance   int64  `json:"balance"`
// }

// func (server *Server) updateAccount(ctx *gin.Context) {
// 	var req updateAccountRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	arg := db.UpdateAccountParams{
// 		ID:       req.ID,
// 		Owner:    req.OwnerName,
// 		Currency: req.Currency,
// 		Balance:  req.Balance,
// 	}

// 	account, err := server.store.UpdateAccount(ctx, arg)

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, account)
// }

// type getAllAccountsRequest struct {
// 	PageNumber int32 `form:"page_number" binding:"required,min=1"`
// 	PerPage    int32 `form:"per_page" binding:"required"`
// }

// func (server *Server) getAllAccounts(ctx *gin.Context) {
// 	var req getAllAccountsRequest
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	arg := db.ListAccountsParams{
// 		Limit:  req.PerPage,
// 		Offset: (req.PageNumber - 1) * req.PerPage,
// 	}

// 	accounts, err := server.store.ListAccounts(ctx, arg)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, gin.H{"accounts": []db.Account{}})
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	if accounts == nil {
// 		ctx.JSON(http.StatusOK, gin.H{"accounts": []db.Account{}})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"accounts": accounts})

// }

// type deleteAccountRequest struct {
// 	ID int64 `uri:"id" binding:"required,min=1"`
// }

// func (server *Server) deleteAccount(ctx *gin.Context) {
// 	var req deleteAccountRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	account, err := server.store.DeleteAccount(ctx, req.ID)
// 	if err != nil {
// 		// if err == sql.ErrNoRows {
// 		// 	ctx.JSON(http.StatusNotFound, errorResponse(err))
// 		// 	return
// 		// }
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, account)

// }
