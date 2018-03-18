package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RootSet ルーターの設定
func RootSet(r *gin.RouterGroup) {
	r.POST("/test", test1)
	r.POST("/test2", test2)
}

func test1(c *gin.Context) {
	var data = []int{1, 2, 3, 4}
	c.JSON(http.StatusOK, data)
}

func test2(c *gin.Context) {
	var data = []int{10, 9, 5, 3}
	c.JSON(http.StatusOK, data)
}
