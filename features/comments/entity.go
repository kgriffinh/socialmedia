package comments

import "github.com/labstack/echo"

type Core struct {
	ID    uint
	Text  string
	Owner string
}

type CommentHandler interface {
	Add() echo.HandlerFunc
	GetComment() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type CommentService interface {
	Add(token interface{}, newComment Core) (Core, error)
	GetComment() ([]Core, error)
	Update(token interface{}, commentID int, updateData Core) (Core, error)
	Delete(token interface{}, commentID int) error
}

type CommentData interface {
	Add(postID int, newPost Core) (Core, error)
	GetComment() ([]Core, error)
	Update(commentID int, postID int, updateData Core) (Core, error)
	Delete(commentID int, postID int) error
}
