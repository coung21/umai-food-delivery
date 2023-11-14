package rest

import (
	"common"
	"net/http"
	"strconv"
	"umai-auth-service/model"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) UpdateRestaurantHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}
		var resUpd model.RestaurantUpdate

		if err := ctx.ShouldBind(&resUpd); err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		resUpdated, err := h.authUC.UpdateRestaurant(ctx, id, &resUpd)
		if err != nil {
			if err == common.NotExistAccount {
				ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
				return
			} else if err == common.Forbidden {
				ctx.JSON(http.StatusForbidden, common.NewRestErr(http.StatusForbidden, err.Error(), err))
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
				return
			}
		}

		ctx.JSON(http.StatusOK, common.NewHttpSuccessResponse(http.StatusOK, "Update restaurant profile successfully.", resUpdated))
	}
}
