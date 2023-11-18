package rest

import (
	"menu-service/middleware"

	"github.com/gin-gonic/gin"
)

func MenuItemRoutes(r *gin.Engine, handlers *menuHandler) {

	authMdw := middleware.Auth(handlers.tokenProvider, handlers.grpcC)
	v1 := r.Group("/v1/restaurant")
	{
		v1.POST("/:id/menu", authMdw, handlers.CreateMenuItemHdl())
		v1.GET("/:id/menu", handlers.ListMenuItemByIDHdl())
		v1.PATCH("/:id/menu/:menu_id", authMdw, handlers.UpdateMenuItem())
		v1.DELETE("/:id/menu/:menu_id", authMdw, handlers.DeleteMenuItemHdl())
	}
}
