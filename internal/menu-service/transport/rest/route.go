package rest

import (
	"menu-service/middleware"

	"github.com/gin-gonic/gin"
)

func MenuItemRoutes(r *gin.Engine, handlers *menuHandler) {

	authMdw := middleware.Auth(handlers.tokenProvider, handlers.grpcC)
	v1 := r.Group("/v1")
	{
		v1.POST("/restaurant/:id/menu", authMdw, handlers.CreateMenuItemHdl())
		v1.GET("/restaurant/:id/menu", handlers.ListMenuItemByIDHdl())
		v1.PATCH("/restaurant/:id/menu/:menu_id", authMdw, handlers.UpdateMenuItem())
		v1.DELETE("/restaurant/:id/menu/:menu_id", authMdw, handlers.DeleteMenuItemHdl())
		v1.GET("/menu/:menu_id", handlers.GetMenuItemHdl())
	}
}
