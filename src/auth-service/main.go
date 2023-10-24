package main

import (
	"umai-auth-service/db"
	"umai-auth-service/middleware"
	"umai-auth-service/transport/rest"

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
	_, err := db.MysqlConn() //db dependency

	if err != nil {
		panic(err)
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
