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
		v1.POST("order/cart", authMdw, handlers.ModifyCartHdl())
		v1.POST("order/cart/delete", authMdw, handlers.DeleteItemFromCartHdl())
	}
}
