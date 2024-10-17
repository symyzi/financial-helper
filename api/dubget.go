package api

// import (
// 	"errors"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	db "github.com/symyzi/financial-helper/db/gen"
// 	"github.com/symyzi/financial-helper/token"
// )

// type budgetCreateRequest struct {
// 	WalletID   int64 `json:"wallet_id"`
// 	Amount     int64 `json:"amount"`
// 	CategoryID int64 `json:"category_id"`
// }

// func (server *Server) createBudget(ctx *gin.Context) {
// 	var req budgetCreateRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	authPayLoad := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
// 	wallet, err := server.store.GetWallet(ctx, req.WalletID)
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, errorResponse(err))
// 		return
// 	}
// 	if wallet.Owner != authPayLoad.Username {
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("unauthorized")))
// 		return
// 	}

// 	arg := db.CreateBudgetParams{
// 		WalletID:   req.WalletID,
// 		Amount:     req.Amount,
// 		CategoryID: req.CategoryID,
// 	}

// 	budget, err := server.store.CreateBudget(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, budget)
// }

// type budgetDeleteRequest struct {
// 	ID       int64 `uri:"id" binding:"required,min=1"`
// 	WalletID int64 `uri:"wallet_id" binding:"required,min=1"`
// }

// // TODO: add delete budget
