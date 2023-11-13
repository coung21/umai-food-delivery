package rest

import (
	"common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) GetRestaurantHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, common.BadQueryParams.Error(), err))
			return
		}

		res, err := h.authUC.GetRestaurant(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		ctx.JSON(http.StatusOK, common.NewHttpSuccessResponse(http.StatusOK, "Get restaurant data successfully.", res))
	}
}
