package middleware

import (
	"common"
	"net/http"
	jwt "order-service/component"
	"order-service/transport/grpc"
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
		// id, err := strconv.Atoi(ctx.Param("id"))
		// if err != nil {
		// 	ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
		// 	ctx.Abort()
		// 	return
		// }
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

		uid, err := grpc.GetUserIdentityHdl(grpcCServ.AuthC, claims.ID)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, common.NewRestErr(http.StatusUnauthorized, common.Unauthorized.Error(), err))
			ctx.Abort()
			return
		}

		// fmt.Println(identity["user_id"])

		if claims.ID == *uid {
			ctx.Set(common.CuserId, *uid)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusForbidden, common.NewRestErr(http.StatusForbidden, common.Forbidden.Error(), err))
			ctx.Abort()
			return
		}
	}
}
