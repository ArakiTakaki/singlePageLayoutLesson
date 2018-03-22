package router

import (
	"github.com/ArakiTakaki/singlePageLayoutLesson/controllers/api"
	"github.com/gin-gonic/gin"
)

// GetRouter そのまま
func GetRouter() *gin.Engine {
	var r = gin.Default()

	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")
	r.Static("/image", "./public/image")
	r.Static("/template", "./public/template")

	r.LoadHTMLGlob("views/*")

	r.NoRoute(func(c *gin.Context) {
		c.HTML(400, "index.html", nil)
	})

	//ルーティンググループテンプレート
	api.RootSet(r.Group("/api"))

	return r
}
