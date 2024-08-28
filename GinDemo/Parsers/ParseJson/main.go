package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func parseJson(c *gin.Context) {
	b, _ := c.GetRawData() // c.Request.Body读取请求数据
	// 定义map或者结构体
	var m map[string]interface{}
	// 反序列化
	_ = json.Unmarshal(b, &m)
	c.JSON(http.StatusOK, m)
}

func main() {
	r := gin.Default()
	r.POST("/parse/json", parseJson)
	r.Run()
}
