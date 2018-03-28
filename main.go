package main

import (
	"github.com/ArakiTakaki/singlePageLayoutLesson/models"
	"github.com/ArakiTakaki/singlePageLayoutLesson/router"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	go def()
	var r = router.GetRouter()
	r.Run(":8000")
}

func def() {
	db := models.Get()
	defer db.Close()
	models.GetAuthor()
}
