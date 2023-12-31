package main

import (
	"fmt"
	jwt "menu-service/component"
	"menu-service/db"
	"menu-service/middleware"
	"menu-service/model"
	mysql_repo "menu-service/repository/mysql_repos"
	redis_repo "menu-service/repository/redis_repos"
	"menu-service/transport/grpc"
	"menu-service/transport/rest"
	"menu-service/usecase"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	// client, err := db.MongoConn()
	db, err := db.MysqlConn()
	//auto migrate
	db.AutoMigrate(&model.MenuItem{})

	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	go func() {
		grpcClient := grpc.RunGrpcClient()
		s.grpcClient = *grpcClient
	}()

	// coll := client.Database(os.Getenv("DB_NAME")).Collection("menus")
	//create index
	// db.CreateIndexes(coll)

	tokenPro := jwt.NewJWTProvider(os.Getenv("SECRET_KEY"))
	menuRepo := mysql_repo.NewMenuRepo(db)
	cacheRepo := redis_repo.NewCacheRepo(fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), os.Getenv("REDIS_PASS"), 0)
	menuUc := usecase.NewMenuUC(menuRepo, cacheRepo)
	menuHdl := rest.NewMenuHandler(menuUc, &s.grpcClient, tokenPro)

	go func() {
		grpc.RunGrpcServer(menuRepo)
	}()

	rest.MenuItemRoutes(r, menuHdl)
	r.Run(":3002")
}
