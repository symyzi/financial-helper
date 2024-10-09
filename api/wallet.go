package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/symyzi/financial-helper/db/gen"
)

type createWalletRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=RUB USD EUR"`
}

func (server *Server) createWallet(ctx *gin.Context) {
	var req createWalletRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateWalletParams{
		Owner:    req.Owner,
		Name:     req.Name,
		Currency: req.Currency,
	}

	wallet, err := server.store.CreateWallet(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, wallet)
}

type getWalletRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getWallet(ctx *gin.Context) {
	var req getWalletRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	wallet, err := server.store.GetWallet(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, wallet)
}

type listWalletsRequest struct {
	Owner    string `form:"owner" binding:"required"`
	PageID   int32  `form:"page_id" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listWallets(ctx *gin.Context) {
	var req listWalletsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListWalletsParams{
		Owner:  req.Owner,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	wallets, err := server.store.ListWallets(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, wallets)
}

type deleteWalletRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteWallet(ctx *gin.Context) {
	var req deleteWalletRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteWallet(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)

}
