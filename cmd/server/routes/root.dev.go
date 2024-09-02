//go:build dev

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpRootGET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func HttpNoRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	}
}
