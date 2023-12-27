package middleware

import (
	"common"
	"fmt"
	jwt "menu-service/component"
	"menu-service/transport/grpc"
	"net/http"
	"strconv"
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

func RestaurantAuth(tokenprovider jwt.TokenProvider, grpcCServ *grpc.GrpcClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			ctx.Abort()
			return
		}
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

		identity, err := grpc.GetResIdentityHdl(grpcCServ.AuthC, claims.ID)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, common.NewRestErr(http.StatusUnauthorized, common.Unauthorized.Error(), err))
			ctx.Abort()
			return
		}

		fmt.Println(identity["user_id"])

		if claims.ID == identity["user_id"] && claims.Role == common.RoleRestaurant && identity["role"] == common.RoleRestaurant && identity["restaurant_id"] == id {
			ctx.Set(common.CuserId, identity["user_id"])
			ctx.Set(common.CuserRole, identity["role"])
			ctx.Set(common.CResId, identity["restaurant_id"])
			ctx.Next()
		} else {
			ctx.JSON(http.StatusForbidden, common.NewRestErr(http.StatusForbidden, common.Forbidden.Error(), err))
			ctx.Abort()
			return
		}
	}
}

func CustomerAuth(tokenprovider jwt.TokenProvider, grpcCServ *grpc.GrpcClient) gin.HandlerFunc {
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

		identity, err := grpc.GetUserIdentityHdl(grpcCServ.AuthC, claims.ID)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, common.NewRestErr(http.StatusUnauthorized, common.Unauthorized.Error(), err))
			ctx.Abort()
			return
		}
		// fmt.Println("user_id", *identity)
		// fmt.Println("user", *identity, claims.ID, claims.Role)
		if claims.ID == *identity {
			ctx.Set(common.CuserId, claims.ID)
			ctx.Set(common.CuserRole, claims.Role)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusForbidden, common.NewRestErr(http.StatusForbidden, common.Forbidden.Error(), err))
			ctx.Abort()
			return
		}
	}
}
