package api

import (
	db "github.com/emonoid/islami_bank_go_backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer (store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

   router.POST("/accounts", server.createAccount)
   router.GET("/accounts/:id", server.getAccount)
   router.GET("/accounts", server.getAllAccounts) 
   router.PUT("accounts/update", server.updateAccount)
   router.DELETE("accounts/delete/:id", server.deleteAccount)

   server.router = router
   return server	
}


func (server *Server) Start(address string) error {
  return  server.router.Run(address)
}


func errorResponse(err error) gin.H {
	return  gin.H{"error": err.Error()}
}