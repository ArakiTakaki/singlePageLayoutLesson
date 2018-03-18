package router

import (
	"net/http"

	context "github.com/ArakiTakaki/Context"
	"github.com/gin-gonic/gin"
)

// GetRouter そのまま
func GetRouter() *gin.Engine {
	var r = gin.Default()
	context.New()

	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")
	r.Static("/image", "./public/image")
	r.Static("/template", "./public/template")

	r.LoadHTMLGlob("views/*")

	r.NoRoute(func(c *gin.Context) {
		c.HTML(400, "index.html", nil)
	})

	r.POST("/api/test", test1)
	r.POST("/api/test2", test2)
	//ルーティンググループテンプレート

	return r
}

func test1(c *gin.Context) {
	var data = []int{1, 2, 3, 4}
	c.JSON(http.StatusOK, data)
}
func test2(c *gin.Context) {
	var data = []int{10, 9, 5, 3}
	c.JSON(http.StatusOK, data)
}
