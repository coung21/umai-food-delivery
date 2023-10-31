package rest

import "github.com/gin-gonic/gin"

func AuthRoutes(r *gin.Engine, handlers *authHandler) {
	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		{

			auth.POST("/register", handlers.RegisterHdl())

		}
	}
}
