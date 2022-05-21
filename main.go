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
		fmt.Println("get server start")
		c.JSON(http.StatusOK, gin.H{
			"img":   controller.GetImageData(c),
			"title": controller.GetTitleData(c),
		})

	})
	r.GET("/set_data", func(c *gin.Context) {
		fmt.Println("set server start")
		img, _ := c.GetQuery("img")
		title, _ := c.GetQuery("title")

		c.JSON(http.StatusOK, gin.H{
			"img":   controller.SetImageData(c, img),
			"title": controller.SetTitleData(c, title),
		})

	})
	r.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
