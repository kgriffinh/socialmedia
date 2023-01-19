package data

import (
	"socialmedia/features/comments"
	"socialmedia/features/comments/data"
	"socialmedia/features/posts"

	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Content     string          `json:"content"`
	Img_content string          `json:"img_content"`
	UserID      uint            `json:"user_id"`
	Comment     []data.Comments `gorm:"foreignKey:PostID"`
}

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Userpp   string
}

type Comment struct {
	ID       uint
	Text     string
	Username string
}

// type File struct {
// 	File multipart.File `json:"file,omitempty"`
// }

func ToCore(data Posts) posts.Core {
	return posts.Core{
		ID:          data.ID,
		Content:     data.Content,
		Img_content: data.Img_content,
	}
}

func CoreToData(data posts.Core) Posts {
	return Posts{
		Model:       gorm.Model{ID: data.ID},
		Content:     data.Content,
		Img_content: data.Img_content,
	}
}

func ToCoreCom(data data.Comments) comments.Core {
	return comments.Core{
		ID:   data.ID,
		Text: data.Text,
	}
}
