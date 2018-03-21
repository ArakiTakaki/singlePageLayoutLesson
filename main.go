package main

import (
	"github.com/ArakiTakaki/singlePageLayoutLesson/router"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	var r = router.GetRouter()
	r.Run(":8000")
}
