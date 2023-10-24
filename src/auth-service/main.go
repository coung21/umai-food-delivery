package main

import (
	"log"
	"umai-auth-service/db"
	"umai-auth-service/transport/rest"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Server struct {
}

func (s *Server) Init(r *gin.Engine) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	_, err := db.MysqlConn()

	if err != nil {
		log.Fatal(err)
	}

	rest.AuthRoutes(r)
}

func main() {
	r := gin.Default()
	var server Server

	server.Init(r)

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":3000")
}
