package api

import (
	"fmt"
	"net/http"

	"github.com/ArakiTakaki/singlePageLayoutLesson/models"

	"github.com/gin-gonic/gin"
)

// RootSet ルーターの設定
func RootSet(r *gin.RouterGroup) {
	r.POST("/test", test1)
	r.POST("/test2", test2)
	r.POST("/author", author)
}

func author(c *gin.Context) {
	work := models.GetAuthor()
	fmt.Println("GET AUTHOR")
	fmt.Println(work)
	c.JSON(http.StatusOK, work)
}

func test1(c *gin.Context) {
	var data = []int{1, 2, 3, 4}
	c.JSON(http.StatusOK, data)
}

func test2(c *gin.Context) {
	var data = []int{10, 9, 5, 3}
	c.JSON(http.StatusOK, data)
}
