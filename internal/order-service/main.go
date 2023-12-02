package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Server struct {
}

func (s *Server) Init(r *gin.Engine) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	var server Server

	server.Init(r)

	r.Run(":3003")
}
