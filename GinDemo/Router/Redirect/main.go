package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/redirect/baidu", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	router.GET("/redirect/:path", func(c *gin.Context) {
		path := c.Param("path")
		c.Request.URL.Path = fmt.Sprintf("/%s", path)
		router.HandleContext(c)
	})

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "you have been redirect internal to /test!",
		})
	})
	router.Run()
}
