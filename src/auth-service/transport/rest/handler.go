package rest

import (
	"net/http"
	"umai-auth-service/common"
	auth "umai-auth-service/interfaces"
	"umai-auth-service/model"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authUC auth.Usecase
}

func NewAuthHandler(authUC auth.Usecase) *authHandler {
	return &authHandler{authUC: authUC}
}

func (h *authHandler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User

		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		//call biz
		newUser, err := h.authUC.Register(&user)
		if err != nil {
			if err.Error() == common.ExistsEmailError.Error() {
				ctx.JSON(http.StatusConflict, common.NewRestErr(http.StatusConflict, err.Error(), err))
				return
			} else if err.Error() == common.BadRequest.Error() {
				ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
			}
		}
		ctx.JSON(http.StatusCreated, newUser)
	}
}
