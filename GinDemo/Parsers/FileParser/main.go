package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// multipart forms默认提交的文件大小限制 32MiB
	// router.MaxMultipartMemory = 8 << 20 // 8 * 2 ^ 20 = 8 MiB

	router.POST("/upload/file", func(c *gin.Context) {
		file, err := c.FormFile("ExampleFile")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		log.Println(file.Filename)
		dst := fmt.Sprintf("./UploadFile/%s", file.Filename)

		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})

	})

	router.POST("/upload/files", func(c *gin.Context) {
		form, err := c.MultipartForm()
		files := form.File["ExampleFiles"]

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("./UploadFiles/%s_%d", file.Filename, index)
			c.SaveUploadedFile(file, dst)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	router.Run()
}
