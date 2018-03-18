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

	r.POST("/api/test", post)
	r.GET("/api/test", post)
	//ルーティンググループテンプレート

	return r
}

func post(c *gin.Context) {
	var data = []int{1, 2, 3, 4}
	c.JSON(http.StatusOK, data)
}
