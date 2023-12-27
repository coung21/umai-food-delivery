package rest

import (
	"common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *menuHandler) ListFavoritesHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. get user id from context
		uid := ctx.Value(common.CuserId).(int)
		// 2. call usecase
		favorites, err := h.menuUC.ListFavorites(ctx, uid)
		if err != nil {
			if err == common.NotFound {
				ctx.JSON(http.StatusNotFound, common.NewRestErr(http.StatusNotFound, err.Error(), err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
			return
		}
		// 3. return favorites
		ctx.JSON(http.StatusOK, common.NewHttpSuccessResponse(http.StatusOK, "success", favorites))
	}
}
