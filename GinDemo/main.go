package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getNormalJson(c *gin.Context) {
	data := map[string]interface{}{
		"name":    "pray",
		"message": "hello.gin!",
		"age":     23,
	}
	c.JSON(http.StatusOK, data)
}

func getStructJson(c *gin.Context) {
	type msg struct {
		Name    string `json:"name"`
		Message string
		Age     int
	}
	data := msg{
		Name:    "小王子",
		Message: "Hello golang!",
		Age:     12,
	}
	c.JSON(http.StatusOK, data)
}

func main() {
	r := gin.Default()
	r.GET("/json/normal/", getNormalJson)
	r.GET("/json/struct/", getStructJson)
)
}
