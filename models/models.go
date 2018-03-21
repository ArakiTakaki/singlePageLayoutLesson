package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model
	ID        uint       `gorm:"primary_key" json: "id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type Post struct {
	Model
	Author   Author
	Title    string `gorm:"size:50" json:title`
	Content  string
	Tags     []Tag `gorm:many2many:post_tags`
	Comments []Comment
}

// Tag メモリに直接記載しても良い
type Tag struct {
	id uint `gorm:AUTO_INCREMENT;PRIMARY_KEY`
}

type Author struct {
	Model
	Name  string `json:"name"`
	Email string `gorm:"size:50" json:"email"`
	URL   string `gorm:"size:50" json:"url"`
}

type Comment struct {
	Model
	Name    string `json:"name"`
	URL     string `json:"url"`
	Content string `json:"comment"`
	Status  bool   `json:"-"`
}
