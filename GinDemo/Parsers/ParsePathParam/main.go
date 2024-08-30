package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func parsePathParam(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(http.StatusOK, gin.H{
		"message":  "ok",
		"username": name,
		"address":  age,
	})
}
func main() {
	r := gin.Default()
	r.GET("/parse/pathparam/:name/:age", parsePathParam)
	r.Run()
}
