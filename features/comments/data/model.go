package data

import (
	"socialmedia/features/comments"

	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	Text   string `json:"text"`
	UserID uint   `json:"user_id"`
	PostID uint   `json:"post_id"`
}

type CommentNonGorm struct {
	ID       uint   `json:"id"`
	Text     string `json:"text"`
	Username string `json:"username"`
	PostID   uint   `json:"post_id"`
}

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

func ToCore(data Comments) comments.Core {
	return comments.Core{
		ID:   data.ID,
		Text: data.Text,
	}
}

func CoreToData(data comments.Core) Comments {
	return Comments{
		Model:  gorm.Model{ID: data.ID},
		Text:   data.Text,
		PostID: data.PostID,
	}
}
