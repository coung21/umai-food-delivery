package rest

import (
	"common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) GetProfileHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		user, err := h.authUC.GetProfile(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		ctx.JSON(http.StatusOK, common.NewHttpSuccessResponse(http.StatusOK, "Get user data successfully.", user))
	}
}
