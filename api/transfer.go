package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	db "github.com/emonoid/islami_bank_go_backend/db/sqlc"
	"github.com/emonoid/islami_bank_go_backend/token"

	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) transferBalance(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	protectedPayload := ctx.MustGet(authorizationPayloadkey).(*token.Payload)

	account, valid := server.validateAccount(ctx, req.FromAccountID, req.Currency)
	if !valid {
		err := errors.New("from account not found")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if account.Owner != protectedPayload.Username { 
		return
	}

	_, valid = server.validateAccount(ctx, req.ToAccountID, req.Currency)

	if !valid { 
		return
	}

	arg := db.TransferTrxnParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTrxn(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validateAccount(ctx *gin.Context, accountID int64, currency string) (db.Account, bool) {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account, false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}

	if account.Currency != currency {
		err := fmt.Errorf("currency doesn't match account [%d] and currency [%s] == [%s]", accountID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return account, false
	}

	return account, true
}
