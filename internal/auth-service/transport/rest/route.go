package rest

import "github.com/gin-gonic/gin"

func AuthRoutes(r *gin.Engine, handlers *authHandler) {
	v1 := r.Group("/v1/auth")
	{
		customer := v1.Group("/customer")
		{

			customer.POST("/register", handlers.RegisterHdl())
			customer.POST("/login", handlers.LoginHdl())
			customer.PATCH("/update/:id", handlers.UpdateUserHdl())
		}

		restaurant := v1.Group("/restaurant")
		{
			restaurant.POST("/register", handlers.RestaurantRegisHdl())
		}
	}
}
