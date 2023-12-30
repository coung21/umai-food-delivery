package rest

import (
	"common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *menuHandler) DeleteFavoriteHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. get user id from context
		uid := ctx.Value(common.CuserId).(int)
		// 2. get menu id from path
		mid, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}
		// 3. call usecase
		favoriteID, err := h.menuUC.DeleteFavorite(ctx, uid, mid)
		if favoriteID == nil && err != nil {
			if err == common.NotFound {
				ctx.JSON(http.StatusNotFound, common.NewRestErr(http.StatusNotFound, err.Error(), err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
			return
		}
		// 4. return favorite id
		ctx.JSON(http.StatusOK, common.NewHttpSuccessResponse(http.StatusOK, "success delete", *favoriteID))
	}
}
