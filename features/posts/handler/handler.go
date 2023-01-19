package handler

import (
	"log"
	"net/http"
	"socialmedia/features/posts"
	"socialmedia/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type postHandle struct {
	srv posts.PostService
}

func New(ps posts.PostService) posts.PostHandler {
	return &postHandle{
		srv: ps,
	}
}

func (ph *postHandle) Add() echo.HandlerFunc {
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

		input := PostRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		cnv := ConvToCore(input)

		res, err := ph.srv.Add(c.Get("user"), *formHeader, *cnv)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessResponse(http.StatusCreated, "sukses menambahkan content", AddToResponse(res)))
	}
}

func (ph *postHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		ParamID := c.Param("post_id")
		postID, _ := strconv.Atoi(ParamID)
		input := PostRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		cnv := ConvToCore(input)
		res, err := ph.srv.Update(c.Get("user"), postID, *cnv) // c.Get("user") user <- table "user" yang diambil adalah ID nya
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessResponse(http.StatusOK, "sukses update content", AddToResponse(res)))
	}
}

func (ph *postHandle) GetPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ph.srv.GetPost()
		if err != nil {
			log.Println("no content found ", err.Error())
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(PrintSuccessResponse(http.StatusOK, "sukses melihat content", GetPostResponse(res)))
	}
}

func (ph *postHandle) GetPostDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		ParamID := c.Param("post_id")
		postID, _ := strconv.Atoi(ParamID)
		res, err := ph.srv.GetPostDetail(postID)
		if err != nil {
			log.Println("no content found ", err.Error())
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(PrintSuccessResponse(http.StatusOK, "sukses melihat content", res))
	}
}

func (ph *postHandle) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		ParamID := c.Param("post_id")
		postID, _ := strconv.Atoi(ParamID)
		err := ph.srv.Delete(c.Get("user"), postID)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessResponse(http.StatusOK, "sukses menghapus content", err))
	}
}
