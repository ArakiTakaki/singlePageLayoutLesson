package router

import (
	context "github.com/ArakiTakaki/Context"
	"github.com/gin-gonic/gin"
)

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

	//ルーティンググループテンプレート

	return r
}
