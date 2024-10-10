package api

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type createCategoryRequest struct {
// 	NameCategory string `json:"category" binding:"required"`
// }

// func (server *Server) createCategory(ctx *gin.Context) {
// 	var req createCategoryRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	category, err := server.store.CreateCategory(ctx, req.NameCategory)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, category)
// }
