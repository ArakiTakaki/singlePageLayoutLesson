package main

import (
	"github.com/ArakiTakaki/singlePageLayoutLesson/router"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// db := models.Get()
	// defer db.Close()
	// models.CreateTables()
	// models.GetAuthor()

	var r = router.GetRouter()
	r.Run(":8000")
}
