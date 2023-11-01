package rest

import (
	"common"
	"net/http"
	"umai-auth-service/model"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) LoginHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var cren model.LoginCredentials

		if err := ctx.ShouldBind(&cren); err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		user, err := h.authUC.Login(ctx.Request.Context(), &cren)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		ctx.JSON(http.StatusOK, common.NewHttpSuccessResponse(http.StatusOK, "Log in successfully.", user))
	}

}
