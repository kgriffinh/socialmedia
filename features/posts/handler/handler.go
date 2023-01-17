package handler

import (
	"log"
	"net/http"
	"socialmedia/features/posts"
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

// Add implements book.BookHandler
func (ph *postHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := PostRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		cnv := ConvToCore(input)

		res, err := ph.srv.Add(c.Get("user"), *cnv) // c.Get("user") user <- table "user" yang diambil adalah ID nya
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessResponse(http.StatusCreated, "sukses menambahkan post", res))
	}
}

// Update implements book.postHandler
func (ph *postHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		ParamBookID := c.Param("post_id")
		postID, _ := strconv.Atoi(ParamBookID)
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
		return c.JSON(PrintSuccessResponse(http.StatusOK, "sukses update post", res))
	}
}

// BookList implements book.postHandler
func (ph *postHandle) GetPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ph.srv.GetPost()
		if err != nil {
			log.Println("no post found ", err.Error())
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		// log.Println(res)
		result := []PostResponse{}
		for i := 0; i < len(res); i++ {
			result = append(result, ToResponse(res[i]))
		}
		return c.JSON(PrintSuccessResponse(http.StatusOK, "sukses melihat post", res))
	}
}

// Delete implements book.postHandler
func (ph *postHandle) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		ParamBookID := c.Param("post_id")
		postID, _ := strconv.Atoi(ParamBookID)
		err := ph.srv.Delete(c.Get("user"), postID)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessResponse(http.StatusOK, "sukses menghapus post", err))
	}
}
