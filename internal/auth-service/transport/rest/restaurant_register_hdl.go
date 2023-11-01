package rest

import (
	"common"
	"net/http"
	"umai-auth-service/model"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) RestaurantRegisHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var restaurant model.Restaurant

		if err := ctx.ShouldBind(&restaurant); err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		res, err := h.authUC.RestaurantRegis(ctx.Request.Context(), &restaurant)
		if err != nil {
			if err == common.NotExistAccout {
				ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
				return
			}
		}

		ctx.JSON(http.StatusCreated, common.NewHttpSuccessResponse(http.StatusCreated, "Register restaurant successfully.", res))
	}
}
