package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func parseFormData(c *gin.Context) {
	username := c.PostForm("username")
	address := c.PostForm("address")

	c.JSON(http.StatusOK, gin.H{
		"message":  "ok",
		"username": username,
		"address":  address,
	})
}

func main() {
	r := gin.Default()
	r.POST("/user/upload", parseFormData)
	r.Run()
}
