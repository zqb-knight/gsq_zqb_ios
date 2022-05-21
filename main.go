package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zqb-knight/gsq_zqb_ios/controller"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/get_data", func(c *gin.Context) {
		fmt.Println("server start")
		c.JSON(http.StatusOK, gin.H{
			"img":   controller.GetImageData(c),
			"title": controller.GetTitleData(c),
		})

	})
	r.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
