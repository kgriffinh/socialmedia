package posts

import "github.com/labstack/echo/v4"

type Core struct {
	ID          uint
	Content     string
	Img_content string
	Owner       string
}

type PostHandler interface {
	Add() echo.HandlerFunc
	GetPost() echo.HandlerFunc
	GetPostDetail() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type PostService interface {
	Add(token interface{}, newPost Core) (Core, error)
	GetPost() ([]Core, error)
	GetPostDetail(postID int) (interface{}, error)
	Update(token interface{}, postID int, updateData Core) (Core, error)
	Delete(token interface{}, postID int) error
}

type PostData interface {
	Add(userID int, newPost Core) (Core, error)
	GetPost() ([]Core, error)
	GetPostDetail(postID int) (interface{}, error)
	Update(postID int, userID int, updateData Core) (Core, error)
	Delete(postID int, userID int) error
}
