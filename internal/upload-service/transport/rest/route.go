package rest

import "github.com/gin-gonic/gin"

func UploadRoutes(r *gin.Engine, handlers *uploadHdl) {
	v1 := r.Group("/v1")
	{
		v1.POST("/upload", handlers.UploadHdl())
	}
}
