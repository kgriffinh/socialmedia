package data

import (
	// "socialmedia/features/posts/data"
	"socialmedia/features/posts/data"
	"socialmedia/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Post     []data.Posts
}

func ToCore(data User) users.Core {
	return users.Core{
		ID:       data.ID,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}
}

func CoreToData(data users.Core) User {
	return User{
		Model:    gorm.Model{ID: data.ID},
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}
}
