package rest

import (
	"net/http"
	"umai-auth-service/common"
	"umai-auth-service/model"

	"github.com/gin-gonic/gin"
)

func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User

		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
		}
	}
}
