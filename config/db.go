package config

import (
	"fmt"
	"log"
	comments "socialmedia/features/comments/data"
	posts "socialmedia/features/posts/data"
	users "socialmedia/features/users/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(ac AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ac.DBUser, ac.DBPass, ac.DBHost, ac.DBPort, ac.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection error : ", err.Error())
		return nil
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&posts.Posts{})
	db.AutoMigrate(&comments.Comments{})

}
