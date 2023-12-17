package rest

import (
	"order-service/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine, handlers *orderHandler) {
	authMdw := middleware.Auth(handlers.tokenPro, handlers.grpcC)
	v1 := r.Group("/v1")
	{
		v1.GET("/order/cart", authMdw, handlers.ListCartItemsHdl())
	}
}
