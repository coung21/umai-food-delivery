package rest

import "github.com/gin-gonic/gin"

func OrderRoutes(r *gin.Engine) {
	r.POST("/order/cart")
}
