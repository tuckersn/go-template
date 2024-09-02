package api

import (
	"github.com/gin-gonic/gin"
)

func HttpProfileIdGET() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, User!",
		})
	}
}
