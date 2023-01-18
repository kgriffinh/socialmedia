package handler

import (
	"net/http"
	"socialmedia/features/comments"
	"strings"
)

type AddCommentResponse struct {
	ID   uint   `json:"id"`
	Text string `json:"comment"`
}

type UpdateCommentResponse struct {
	Text string `json:"comment"`
}

func AddToResponse(data comments.Core) AddCommentResponse {
	return AddCommentResponse{
		ID:   data.ID,
		Text: data.Text,
	}
}

func UpdToResponse(data comments.Core) UpdateCommentResponse {
	return UpdateCommentResponse{
		Text: data.Text,
	}
}

func PrintSuccessResponse(code int, message string, data ...interface{}) (int, interface{}) {
	resp := map[string]interface{}{}
	if len(data) < 2 {
		resp["data"] = (data[0])
	} else {
		resp["data"] = (data[0])
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
