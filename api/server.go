package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/symyzi/financial-helper/db/gen"
	"github.com/symyzi/financial-helper/token"
	"github.com/symyzi/financial-helper/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/")
	authRoutes.Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/wallets", server.createWallet)
	authRoutes.GET("/wallets", server.listWallets)
	authRoutes.GET("/wallets/:id", server.getWallet)
	authRoutes.DELETE("/wallets/:id", server.deleteWallet)

	walletRoutes := authRoutes.Group("/wallets/:id")

	walletRoutes.POST("/expenses", server.createExpense)
	walletRoutes.GET("/expenses", server.listExpenses)
	walletRoutes.GET("/expenses/:id", server.getExpense)
	walletRoutes.DELETE("/expenses/:id", server.deleteExpense)

	// budgetRoutes := walletRoutes.Group("/:id/budgets")
	// budgetRoutes.POST("/", server.createBudget)
	// budgetRoutes.GET("/", server.listBudgets)
	// budgetRoutes.GET("/:id", server.getBudget)
	// budgetRoutes.DELETE("/:id", server.deleteBudget)

	server.router = router
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	if err == nil {
		return gin.H{"error": "unknown error"}
	}
	return gin.H{"error": err.Error()}
}
