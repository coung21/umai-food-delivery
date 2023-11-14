package rest

import (
	"umai-auth-service/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, handlers *authHandler) {

	authMdw := middleware.Auth(handlers.tokenProvider, handlers.authRepo)
	v1 := r.Group("/v1/auth")
	{
		customer := v1.Group("/customer")
		{

			customer.POST("/register", handlers.RegisterHdl())
			customer.POST("/login", handlers.LoginHdl())
			customer.GET("/:id", authMdw, handlers.GetProfileHdl())
			customer.PATCH("/:id", authMdw, handlers.UpdateUserHdl())
		}

		restaurant := v1.Group("/restaurant")
		{
			restaurant.POST("/register", authMdw, handlers.RestaurantRegisHdl())
			restaurant.GET("/:id", handlers.GetRestaurantHdl())
			restaurant.PATCH("/:id", authMdw, handlers.UpdateRestaurantHdl())
		}
	}
}
