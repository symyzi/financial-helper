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

	walletRoutes := authRoutes.Group("/wallets")

	walletRoutes.POST("/", server.createWallet)
	walletRoutes.GET("/", server.listWallets)
	walletRoutes.GET("/:walletID", server.getWallet)
	walletRoutes.DELETE("/:walletID", server.deleteWallet)

	expenseRoutes := walletRoutes.Group("/:walletID/expenses")

	expenseRoutes.POST("/", server.createExpense)
	expenseRoutes.GET("/", server.listExpenses)
	expenseRoutes.GET("/:expenseID", server.getExpense)
	// expenseRoutes.DELETE("/:expenseID", server.deleteExpense)

	// budgetRoutes := walletRoutes.Group("/:walletID/budgets")
	// budgetRoutes.POST("/", server.createBudget)
	// budgetRoutes.GET("/", server.listBudgets)
	// budgetRoutes.GET("/:budgetID", server.getBudget)
	// budgetRoutes.DELETE("/:budgetID", server.deleteBudget)

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
