package data

import (
	c "socialmedia/features/comments/data"
	"socialmedia/features/posts/data"
	"socialmedia/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Userpp   string
	Post     []data.Posts `gorm:"foreignKey:UserID"`
	Comment  []c.Comments `gorm:"foreignKey:UserID"`
}

func ToCore(data User) users.Core {
	return users.Core{
		ID:       data.ID,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Userpp:   data.Userpp,
	}
}

func CoreToData(data users.Core) User {
	return User{
		Model:    gorm.Model{ID: data.ID},
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Userpp:   data.Userpp,
	}
}
