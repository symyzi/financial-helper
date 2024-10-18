package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/symyzi/financial-helper/db/gen"
	"github.com/symyzi/financial-helper/token"
)

type budgetCreateRequest struct {
	WalletID   int64 `json:"wallet_id"`
	Amount     int64 `json:"amount"`
	CategoryID int64 `json:"category_id"`
}

func (server *Server) createBudget(ctx *gin.Context) {
	var req budgetCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayLoad := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	wallet, err := server.store.GetWallet(ctx, req.WalletID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	if wallet.Owner != authPayLoad.Username {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("unauthorized")))
		return
	}

	arg := db.CreateBudgetParams{
		WalletID:   req.WalletID,
		Amount:     req.Amount,
		CategoryID: req.CategoryID,
	}

	budget, err := server.store.CreateBudget(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, budget)
}

type budgetDeleteRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteBudget(ctx *gin.Context) {
	var req budgetDeleteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayLoad := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	budget, err := server.store.GetBudgetByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(
			http.StatusInternalServerError,
			errorResponse(err),
		)
		return
	}

	wallet, err := server.store.GetWallet(ctx, budget.WalletID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	if wallet.Owner != authPayLoad.Username {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("unauthorized")))
		return
	}

	err = server.store.DeleteBudget(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

type budgetGetRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getBudget(ctx *gin.Context) {
	var req budgetGetRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayLoad := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	budget, err := server.store.GetBudgetByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(
			http.StatusInternalServerError,
			errorResponse(err),
		)
		return
	}

	wallet, err := server.store.GetWallet(ctx, budget.WalletID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	if wallet.Owner != authPayLoad.Username {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("unauthorized")))
		return
	}
	ctx.JSON(http.StatusOK, budget)
}

type budgetListRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listBudgets(ctx *gin.Context) {
	var req budgetListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayLoad := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.ListBudgetsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	budgets, err := server.store.ListBudgets(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for _, budget := range budgets {
		wallet, err := server.store.GetWallet(ctx, budget.WalletID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		if wallet.Owner != authPayLoad.Username {
			ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("unauthorized")))
			return
		}
	}
	ctx.JSON(http.StatusOK, budgets)
}
