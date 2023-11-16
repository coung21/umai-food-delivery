package main

import (
	"context"
	"menu-service/middleware"
	"menu-service/transport/grpc"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type server struct {
	grpcClient grpc.GrpcClient
}

func main() {
	s := &server{}

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Use(middleware.Recover())
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	go func() {
		grpcClient := grpc.RunGrpcClient()
		s.grpcClient = *grpcClient
	}()

	// coll := client.Database(os.Getenv("DB_NAME")).Collection("menus")
	// tokenPro := jwt.NewJWTProvider(os.Getenv("SECRET_KEY"))
	// menuRepo := repository.NewMenuRepo(coll)
	// authUc := usecase.NewAuthUC(authRepo, tokenPro, 24*10)
	// authHdl := rest.NewAuthHandler(authUc, authRepo, tokenPro)

	// rest.AuthRoutes(r, authHdl)
	r.Run(":3002")
}
