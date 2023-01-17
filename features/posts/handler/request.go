package handler

import "socialmedia/features/posts"

type PostRequest struct {
	Content     string `json:"content" form:"content"`
	Img_content string `json:"img_content" form:"img_content"`
}

func ConvToCore(data interface{}) *posts.Core {
	res := posts.Core{}

	switch data.(type) {
	case PostRequest:
		cnv := data.(PostRequest)
		res.Content = cnv.Content
		res.Img_content = cnv.Img_content
	default:
		return nil
	}

	return &res
}
