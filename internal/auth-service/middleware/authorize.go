package middleware

import (
	"common"
	"net/http"
	"strings"
	jwt "umai-auth-service/component"
	auth "umai-auth-service/interfaces"

	"github.com/gin-gonic/gin"
)

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", common.InvalidAuthHeader
	}

	return parts[1], nil
}

func Auth(tokenprovider jwt.TokenProvider, authRepo auth.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := extractTokenFromHeaderString(ctx.GetHeader("Authorization"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, common.NewRestErr(http.StatusUnauthorized, err.Error(), err))
			ctx.Abort()
			return
		}

		claims, err := tokenprovider.Validate(token)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, common.NewRestErr(http.StatusUnauthorized, common.InvalidJWTToken.Error(), err))
			ctx.Abort()
			return
		}

		user, err := authRepo.FindUserByID(ctx.Request.Context(), claims.ID)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, common.NewRestErr(http.StatusUnauthorized, common.InvalidJWTToken.Error(), err))
			ctx.Abort()
			return
		}

		ctx.Set(common.CuserId, user.ID)
		ctx.Set(common.CuserRole, user.Role)
		ctx.Next()
	}
}
