package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/symyzi/financial-helper/db/gen"
	"github.com/symyzi/financial-helper/token"
)

type createExpenseRequest struct {
	WalletID           int64          `json:"wallet_id"`
	Amount             int64          `json:"amount"`
	ExpenseDescription sql.NullString `json:"expense_description"`
	CategoryID         int64          `json:"category_id"`
	ExpenseDate        time.Time      `json:"expense_date"`
}

func (server *Server) createExpense(ctx *gin.Context) {

	authPayLoad := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	var req createExpenseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	wallet, err := server.store.GetWallet(ctx, req.WalletID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	if wallet.Owner != authPayLoad.Username {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("unauthorized")))
		return
	}

	arg := db.CreateExpenseParams{
		WalletID:           req.WalletID,
		Amount:             req.Amount,
		ExpenseDescription: req.ExpenseDescription,
		CategoryID:         req.CategoryID,
		ExpenseDate:        req.ExpenseDate,
	}
	expense, err := server.store.CreateExpense(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, expense)
}

type listExpensesRequest struct {
	WalletID int64 `form:"wallet_id" binding:"required"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listExpenses(ctx *gin.Context) {
	var req listExpensesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
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

	arg := db.ListExpensesParams{
		WalletID: req.WalletID,
		Limit:    req.PageSize,
		Offset:   (req.PageID - 1) * req.PageSize,
	}

	expenses, err := server.store.ListExpenses(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, expenses)
}

type getExpenseRequest struct {
	ID       int64 `uri:"id" binding:"required,min=1"`
	WalletID int64 `uri:"wallet_id" binding:"required,min=1"`
}

func (server *Server) getExpense(ctx *gin.Context) {
	var req getExpenseRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
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

	expense, err := server.store.GetExpense(ctx, db.GetExpenseParams{
		ID:       req.ID,
		WalletID: req.WalletID,
	})
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
	ctx.JSON(http.StatusOK, expense)
}

type deleteExpenseRequest struct {
	ID       int64 `uri:"id" binding:"required,min=1"`
	WalletID int64 `uri:"wallet_id" binding:"required,min=1"`
}

// TODO: add delete expense
