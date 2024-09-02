package api

import (
	"github.com/gin-gonic/gin"
)

func HttpUserGET() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, User!",
		})
	}
}

func HttpUserPOST() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "User Created",
		})
	}
}
