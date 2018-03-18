package main

import "github.com/ArakiTakaki/singlePageLayoutLesson/router"

func main() {
	var r = router.GetRouter()
	r.Run(":8000")
}
