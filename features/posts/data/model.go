package data

import (
	// "socialmedia/features/posts/data"
	"socialmedia/features/posts"

	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Content     string
	Img_content string
	UserID      uint
}

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

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
