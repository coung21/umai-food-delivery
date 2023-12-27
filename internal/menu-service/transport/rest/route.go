package rest

import (
	"menu-service/middleware"

	"github.com/gin-gonic/gin"
)

func MenuItemRoutes(r *gin.Engine, handlers *menuHandler) {

	RestaurantAuthMdw := middleware.RestaurantAuth(handlers.tokenProvider, handlers.grpcC)
	CustomerAuthMdw := middleware.CustomerAuth(handlers.tokenProvider, handlers.grpcC)
	v1 := r.Group("/v1")
	{
		v1.POST("/restaurant/:id/menu", RestaurantAuthMdw, handlers.CreateMenuItemHdl())
		v1.GET("/restaurant/:id/menu", handlers.ListMenuItemByIDHdl())
		v1.PATCH("/restaurant/:id/menu/:menu_id", RestaurantAuthMdw, handlers.UpdateMenuItem())
		v1.DELETE("/restaurant/:id/menu/:menu_id", RestaurantAuthMdw, handlers.DeleteMenuItemHdl())
		v1.GET("/menu/:menu_id", handlers.GetMenuItemHdl())
		v1.POST("/menu/:id/favorite", CustomerAuthMdw, handlers.AddFavoriteHdl())
	}
}
