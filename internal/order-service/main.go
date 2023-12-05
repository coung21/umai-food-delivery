package main

import (
	"encoding/json"
	"fmt"
	"order-service/transport/grpc"

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

	// cdb := db.RedisConn(fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), os.Getenv("REDIS_PASS"), 0)

	go func() {
		grpcClient := grpc.RunGrpcClient()
		s.grpcClient = *grpcClient
	}()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{"message": "pong"})
	})

	var server Server

	server.Init(r)

	r.GET("/lol", func(ctx *gin.Context) {

		re, err := grpc.GetMenuItemHdl(server.grpcClient.MenuC, "6558f5358b01596f94ff97f1")
		if err != nil {
			fmt.Println(err.Error())

		}

		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(re), &jsonMap)
		ctx.JSON(200, gin.H{"message": jsonMap})
	})

	r.Run(":3003")
}
