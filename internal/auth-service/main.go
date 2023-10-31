package main

import (
	"os"
	jwt "umai-auth-service/component"
	"umai-auth-service/db"
	"umai-auth-service/middleware"
	"umai-auth-service/repository"
	"umai-auth-service/transport/rest"
	"umai-auth-service/usecase"

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
	db, err := db.MysqlConn()

	if err != nil {
		panic(err)
	}

	tokenPro := jwt.NewJWTProvider(os.Getenv("SECRET_KEY"))
	authRepo := repository.NewAuthRepo(db)
	authUc := usecase.NewAuthUC(authRepo, tokenPro, 24*10)
	authHdl := rest.NewAuthHandler(authUc)

	rest.AuthRoutes(r, authHdl)
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

	r.Run(":3000")
}
