package main

import (
	"fmt"
	"os"
	jwt "umai-auth-service/component"
	"umai-auth-service/db"
	"umai-auth-service/model"
	mysql_repo "umai-auth-service/repository/mysql_repos"
	cache_repo "umai-auth-service/repository/redis_repos"
	"umai-auth-service/transport/grpc"
	"umai-auth-service/transport/rest"
	"umai-auth-service/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Server struct {
}

func (s *Server) Init(r *gin.Engine) {
	// r.Use(middleware.Recover())
	// r.Use(gin.Recovery())
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	mdb, err := db.MysqlConn()
	//auto migrate
	mdb.AutoMigrate(&model.User{})
	mdb.AutoMigrate(&model.Restaurant{})

	//redis conn
	cdb := db.RedisConn(fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), os.Getenv("REDIS_PASS"), 0)
	if err != nil {
		panic(err)
	}

	tokenPro := jwt.NewJWTProvider(os.Getenv("SECRET_KEY"))
	authRepo := mysql_repo.NewAuthRepo(mdb)
	cacheRepo := cache_repo.NewCacheAuthRepo(cdb)
	authUc := usecase.NewAuthUC(authRepo, cacheRepo, tokenPro, 24*10)
	authHdl := rest.NewAuthHandler(authUc, authRepo, tokenPro)

	go func() {
		grpc.RunGrpcServer(authRepo)
	}()

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
