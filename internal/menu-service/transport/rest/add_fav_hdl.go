package rest

import (
	"common"
	"menu-service/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *menuHandler) AddFavoriteHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. get user id from context
		uid := ctx.Value(common.CuserId).(int)
		// 2. get menu id from path
		mid := ctx.Param("id")
		// 3. call usecase
		favoriteID, err := h.menuUC.AddFavorite(ctx, uid, mid)
		if err != nil {
			if err == common.NotFound {
				ctx.JSON(http.StatusNotFound, common.NewRestErr(http.StatusNotFound, err.Error(), err))
				return
			} else if err == model.ErrAlreadyFavorite {
				ctx.JSON(http.StatusConflict, common.NewRestErr(http.StatusConflict, err.Error(), err))
				return
			} else if err == common.BadQueryParams {
				ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
			return
		}
		// 4. return favorite id
		ctx.JSON(http.StatusOK, common.NewHttpSuccessResponse(http.StatusOK, "success", *favoriteID))
	}
}
