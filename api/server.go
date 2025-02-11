package api

import (
	db "github.com.victorex27/simple_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.PATCH("/accounts/:id/amount/:amount", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	server.router = router
	return server
}

// router field is private and cannot be accessed outside of this package
func (server *Server) Start(address string) error{
	return server.router.Run(address)
}


func errorResponse(err error) gin.H{
	return gin.H{"error": err.Error()}
}