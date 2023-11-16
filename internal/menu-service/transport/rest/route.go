package rest

import (
	"menu-service/middleware"

	"github.com/gin-gonic/gin"
)

func MenuItemRoutes(r *gin.Engine, handlers *menuHandler) {

	authMdw := middleware.Auth(handlers.tokenProvider, handlers.grpcC)
	v1 := r.Group("/v1/restaurant/menu")
	{
		v1.POST("/", authMdw, handlers.CreateMenuItemHdl())
	}
}
