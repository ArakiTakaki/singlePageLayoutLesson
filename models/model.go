package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db = newDBConn()

// NewDBConn (DBの接続を返す)
func newDBConn() *gorm.DB {

	const CONNECT = "root@/shingle_lesson"
	db, err := gorm.Open("mysql", CONNECT)
	if err != nil {
		panic(err)
	}
	return db
}

// Get DBの接続を返す（基本的に
func Get() *gorm.DB {
	return db
}

type Place struct {
	gorm.Model
	ID int `gorm:size""`
}

func main() {
	test := Get()
	test.Close()
	test = Get()
	test.Close()
	test = Get()
	test.Close()
}
