package rest

import (
	"common"
	"net/http"
	"strconv"
	"umai-auth-service/model"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) UpdateUserHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		var userUdp model.UserUpdate

		if err := ctx.ShouldBind(&userUdp); err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		userUpdated, err := h.authUC.UpdateProfile(ctx, id, &userUdp)
		if err != nil {
			if err == common.NotExistAccount {
				ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
				return
			}
		}

		ctx.JSON(http.StatusOK, common.NewHttpSuccessResponse(http.StatusOK, "Update user profile successfully.", userUpdated))
	}
}
