package rest

import "github.com/gin-gonic/gin"

func AuthRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			customer := auth.Group("/customer")
			{
				customer.POST("/register")
			}
		}
	}
}
