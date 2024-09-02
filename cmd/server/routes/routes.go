package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpRedirect(redirectPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.Host + redirectPath
		fmt.Println("Redirecting to", path)
		c.Redirect(http.StatusTemporaryRedirect, path)
	}
}
