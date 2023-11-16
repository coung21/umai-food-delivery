package middleware

import (
	"common"
	"fmt"
	jwt "menu-service/component"
	"menu-service/transport/grpc"
	"net/http"
	"strings"

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

func Auth(tokenprovider jwt.TokenProvider, grpcCServ *grpc.GrpcClient) gin.HandlerFunc {
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

		uid, urole := grpc.GetIdentityHdl(grpcCServ.Client, claims.ID)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, common.NewRestErr(http.StatusUnauthorized, common.InvalidJWTToken.Error(), err))
			ctx.Abort()
			return
		}

		fmt.Println(claims.ID, uid)

		if claims.ID == uid && urole == "restaurant" {
			ctx.Set(common.CuserId, uid)
			ctx.Set(common.CuserRole, urole)
			ctx.Next()
		} else {
			ctx.Abort()
		}
	}
}
