package handler

import "socialmedia/features/comments"

type CommentRequest struct {
	PostID uint   `json:"post_id" form:"post_id"`
	Text   string `json:"comment" form:"comment"`
}

func ReqToCore(data interface{}) *comments.Core {
	res := comments.Core{}

	switch data.(type) {
	case CommentRequest:
		cnv := data.(CommentRequest)
		res.PostID = cnv.PostID
		res.Text = cnv.Text
	default:
		return nil
	}

	return &res
}
