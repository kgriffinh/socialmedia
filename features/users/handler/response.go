package handler

import (
	"net/http"
	"socialmedia/features/users"
	"strings"
)

type UserReponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func ToResponse(data users.Core) UserReponse {
	return UserReponse{
		Username: data.Username,
		Email:    data.Email,
	}
}

func PrintSuccessReponse(code int, message string, data ...interface{}) (int, interface{}) {
	resp := map[string]interface{}{}
	if len(data) < 2 {
		resp["data"] = ToResponse(data[0].(users.Core))
	} else {
		resp["data"] = ToResponse(data[0].(users.Core))
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
