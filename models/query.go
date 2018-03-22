package models

import (
	"fmt"

	context "github.com/ArakiTakaki/Context"
)

// GetAuthor 記者一覧を返却するメソッド
func GetAuthor() interface{} {
	value, err := context.Get("Author")
	if err != nil {
		db := Get()
		defer db.Close()
		fmt.Println("DBから記述者情報をゲットしました。")
		query := db.Find(&Author{})
		context.Put("Author", query.Value)
		return query.Value
	}
	return value
}
