package comments

import "github.com/labstack/echo/v4"

type Core struct {
	ID     uint
	Text   string `validate:"required"`
	Owner  string
	PostID uint
}

type CommentHandler interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type CommentService interface {
	Add(token interface{}, newComment Core) (Core, error)
	Update(token interface{}, commentID int, updateData Core) (Core, error)
	Delete(token interface{}, commentID int) error
}

type CommentData interface {
	Add(userID int, newPost Core) (Core, error)
	Update(commentID int, postID int, updateData Core) (Core, error)
	Delete(commentID int, postID int) error
}
