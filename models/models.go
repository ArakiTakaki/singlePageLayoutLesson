package models

import (
	"time"
)

// CreateTables テーブル変更のマイグレーションファイル
func CreateTables() {
	var db = Get()
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Post{})
}

// Model 初期セット
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// Post 記事の投下部
type Post struct {
	Model
	Title    string `gorm:"size:50" json:"title"`
	Content  string
	AuthorID uint
	Author   Author `gorm:"ForeignKey:AuthorID;AssociationForeignKey:ID"`
	Tags     []Tag  `gorm:"many2many:post_tags"`
	Comments []Comment
}

// Tag メモリに直接記載しても良い
type Tag struct {
	ID       uint   `gorm:"AUTO_INCREMENT; PRIMARY_KEY" json:"id"`
	Name     string `gorm:"size:50" json:"name"`
	ParentID uint   `json:"parent"`
}

// Author 記事の筆者情報
type Author struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Email string `gorm:"size:50" json:"email"`
	URL   string `gorm:"size:50" json:"url"`
}

// Comment 記事に対するコメント
type Comment struct {
	Model
	PostID  uint   `gorm:"index"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	Content string `json:"comment"`
	Status  bool   `json:"-"`
}
