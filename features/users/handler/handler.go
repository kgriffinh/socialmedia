package handler

import (
	"net/http"
	"socialmedia/features/users"
	"socialmedia/helper"

	"github.com/labstack/echo/v4"
)

type userControll struct {
	srv users.UserService
}

func New(srv users.UserService) users.UserHandler {
	return &userControll{
		srv: srv,
	}
}

func (uc *userControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		token, res, err := uc.srv.Login(input.Username, input.Password)
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(PrintSuccessReponse(http.StatusOK, "berhasil login", ToResponse(res), token))
	}
}
func (uc *userControll) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		res, err := uc.srv.Register(*ReqToCore(input))
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(PrintSuccessReponse(http.StatusCreated, "berhasil mendaftar", ToResponse(res)))
	}
}
func (uc *userControll) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := uc.srv.Profile(token)
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(PrintSuccessReponse(http.StatusOK, "berhasil lihat profil", PPToResponse(res)))
	}
}

func (uc *userControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		formHeader, err := c.FormFile("file")
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				helper.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       &echo.Map{"data": "Select a file to upload"},
				})
		}

		token := c.Get("user")
		input := UpdateRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		res, err := uc.srv.Update(token, *formHeader, *ReqToCore(input))
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessReponse(http.StatusOK, "berhasil update profil", PPToResponse(res)))
	}
}

func (uc *userControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		err := uc.srv.Delete(token)
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(PrintSuccessReponse(http.StatusOK, "berhasil delete profil", err))
	}
}
