package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getQueryWithParam(c *gin.Context) {
	param := c.Query("param")
	defaultParam := c.DefaultQuery("defaultParam", "This is value of default param.")
	desireParam, ok := c.GetQuery("desireParam")

	if !ok {
		desireParam = "not provided"
	}

	c.JSON(http.StatusOK, gin.H{
		"param":        param,
		"defualtParam": defaultParam,
		"desireParam":  desireParam,
	})
}

func main() {
	r := gin.Default()
	r.GET("/json/param/", getQueryWithParam) // param=hello&defaultParam=xx
	r.Run()
}
