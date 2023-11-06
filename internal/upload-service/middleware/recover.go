package middleware

import (
	"common"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatal(err)
				ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, "Something went wrong with the server", err))
			}
		}()

		ctx.Next()
	}
}
