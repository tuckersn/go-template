package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/tuckersn/go-template/cmd/server/routes"
	"github.com/tuckersn/go-template/cmd/server/routes/api"
	"github.com/tuckersn/go-template/internal/global"
)

func HttpLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Info().Str(
			"method",
			ctx.Request.Method,
		).Str(
			"ip",
			ctx.ClientIP(),
		).Str(
			"path",
			ctx.Request.URL.Path,
		).Msg("HTTP Request")
		ctx.Next()
	}
}

func httpServer() *gin.Engine {
	r := gin.New(func(e *gin.Engine) {
		e.Use(gin.Recovery())
		e.Use(HttpLogger())
	})

	r.GET("/", routes.HttpRootGET)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiRoutes := r.Group("/api")
	{
		hello := apiRoutes.Group("/hello")
		{
			hello.GET("/", api.HttpHelloGET())
		}

		user := apiRoutes.Group("/user")
		{
			user.GET("/", api.HttpUserGET())
			user.POST("/", api.HttpUserPOST())

			userProfile := user.Group("/profile")
			{
				userProfile.GET("/:id", api.HttpProfileIdGET())
			}
		}
	}

	r.NoRoute(routes.HttpNoRoute())

	return r
}

func main() {

	err := global.InitContext()
	if err != nil {
		panic(err)
	}

	{
		portStr, ok := os.LookupEnv("PORT")
		if ok {
			global.API_PORT, err = strconv.Atoi(portStr)
			if err != nil {
				panic(err)
			}
		}
	}

	ginServer := httpServer()
	errChan := make(chan error)

	go func(errChan chan error) {
		fmt.Println("server starting on http://localhost:" + strconv.Itoa(global.API_PORT))
		err := ginServer.Run("0.0.0.0:" + strconv.Itoa(global.API_PORT))
		if err != nil {
			errChan <- err
		} else {
			fmt.Println("server closed unexpectedly")
		}
	}(errChan)

	for {
		err := <-errChan
		fmt.Println(err)
	}
}
