package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/symyzi/financial-helper/db/gen"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)

	router.POST("/wallets", server.createWallet)
	router.GET("/wallets/:id", server.getWallet)
	router.GET("/wallets", server.listWallets)
	router.DELETE("/wallets/:id", server.deleteWallet)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
