package main

import (
	"os"
	"upload-service/component/uploadprovider"
	"upload-service/middleware"
	"upload-service/transport/rest"
	"upload-service/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Server struct {
}

func (s *Server) Init(r *gin.Engine) {
	r.Use(middleware.Recover())
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	s3BucketName := os.Getenv("S3_BUCKET")
	s3Region := os.Getenv("S3_REGION")
	s3AccessKey := os.Getenv("S3_ACCESS_KEY")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")
	s3Domain := os.Getenv("S3_DOMAIN")

	provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3AccessKey, s3SecretKey, s3Domain)
	uc := usecase.NewUploadUC(provider)
	hdl := rest.NewUploadHdl(uc)
	rest.UploadRoutes(r, hdl)
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	var server Server

	server.Init(r)

	r.Run(":3001")
}
