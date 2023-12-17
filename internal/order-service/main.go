package main

import (
	"fmt"
	jwt "order-service/component"
	"order-service/db"
	repository "order-service/repository/redis_repos"
	"order-service/transport/grpc"
	"order-service/transport/rest"
	"order-service/usecase"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Server struct {
	grpcClient grpc.GrpcClient
}

func (s *Server) Init(r *gin.Engine) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cdb := db.RedisConn(fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), os.Getenv("REDIS_PASS"), 0)

	go func() {
		grpcClient := grpc.RunGrpcClient()
		s.grpcClient = *grpcClient
	}()

	tokenPro := jwt.NewJWTProvider(os.Getenv("SECRET_KEY"))
	orderRepo := repository.NewCacheRepo(cdb)
	orderUc := usecase.NewOrderUC(orderRepo)
	orderHdl := rest.NewOrderHandler(orderUc, tokenPro, &s.grpcClient)

	rest.OrderRoutes(r, orderHdl)
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{"message": "pong"})
	})

	var server Server

	server.Init(r)

	// r.GET("/lol", func(ctx *gin.Context) {

	// 	re, err := grpc.GetMenuItemHdl(server.grpcClient.MenuC, "6558f5358b01596f94ff97f1")
	// 	if err != nil {
	// 		fmt.Println(err.Error())

	// 	}

	// 	var jsonMap map[string]interface{}
	// 	json.Unmarshal([]byte(re), &jsonMap)
	// 	ctx.JSON(200, gin.H{"message": jsonMap})
	// })

	r.Run(":3003")
}
