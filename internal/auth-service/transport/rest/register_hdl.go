package rest

import (
	"common"
	"net/http"
	"umai-auth-service/model"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) RegisterHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User

		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}
		//call biz
		newUser, err := h.authUC.Register(ctx.Request.Context(), &user)
		if err != nil {
			if err == common.ExistsEmailError {
				ctx.JSON(http.StatusConflict, common.NewRestErr(http.StatusConflict, err.Error(), err))
				return
			} else if err == common.BadRequest {
				ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
				return
			}
		}
		ctx.JSON(http.StatusCreated, common.NewHttpSuccessResponse(http.StatusCreated, "Register user successfully.", newUser))
	}
}
