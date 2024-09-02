//go:build !dev

package routes

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

//go:embed html/*
var reactFS embed.FS

func serveHtml(c *gin.Context, filePath string) {
	file, err := reactFS.Open(filePath)
	if err != nil {
		c.String(http.StatusNotFound, "File not found")
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not retrieve file information")
		return
	}

	data, err := fs.ReadFile(reactFS, filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not read file")
		return
	}

	contentType := "text/plain"
	switch path.Ext(filePath) {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	}

	c.Header("Content-Type", contentType)
	c.Header("Content-Length", strconv.Itoa(len(data)))
	http.ServeContent(c.Writer, c.Request, fileInfo.Name(), fileInfo.ModTime(), bytes.NewReader(data))
}

func HttpRootGET(c *gin.Context) {

	serveHtml(c, "html/index.html")
}

func HttpNoRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		if _, err := reactFS.Open("html" + c.Request.URL.Path); err == nil {
			fmt.Println("Serving", c.Request.URL.Path)
			serveHtml(c, "html"+c.Request.URL.Path)
			return
		} else {
			fmt.Println("Not Found", c.Request.URL.Path)
		}

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
	}
}
