package handler

import (
	"net/http"
	"socialmedia/features/posts"
	"strings"
)

type PostResponse struct {
	ID          uint   `json:"id"`
	Content     string `json:"content"`
	Img_content string `json:"img_content"`
	Owner       string `json:"owner"`
}

func ToResponse(data posts.Core) PostResponse {
	return PostResponse{
		ID:          data.ID,
		Content:     data.Content,
		Img_content: data.Img_content,
		Owner:       data.Owner,
	}
}

func PrintSuccessResponse(code int, message string, data ...interface{}) (int, interface{}) {
	resp := map[string]interface{}{}
	if len(data) < 2 {
		resp["data"] = PostResponse(data[0].(posts.Core))
	} else {
		resp["data"] = PostResponse(data[0].(posts.Core))
		resp["token"] = data[1].(string)
	}

	if message != "" {
		resp["message"] = message
	}

	return code, resp
}

func PrintErrorResponse(msg string) (int, interface{}) {
	resp := map[string]interface{}{}
	code := -1
	if msg != "" {
		resp["message"] = msg
	}

	if strings.Contains(msg, "server") {
		code = http.StatusInternalServerError
	} else if strings.Contains(msg, "format") {
		code = http.StatusBadRequest
	} else if strings.Contains(msg, "not found") {
		code = http.StatusNotFound
	}

	return code, resp
}
