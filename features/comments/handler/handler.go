package handler

import (
	"log"
	"net/http"
	"socialmedia/features/comments"
	"strconv"

	"github.com/labstack/echo"
)

type commentHandle struct {
	srv comments.CommentService
}

func New(cs comments.CommentService) comments.CommentHandler {
	return &commentHandle{
		srv: cs,
	}
}

func (ch *commentHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := CommentRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		cnv := ReqToCore(input)

		res, err := ch.srv.Add(c.Get("user"), *cnv)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessResponse(http.StatusCreated, "sukses menambahkan comment", AddToResponse(res)))
	}
}

func (ch *commentHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		ParamID := c.Param("comment_id")
		commentID, _ := strconv.Atoi(ParamID)
		input := CommentRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		cnv := ReqToCore(input)
		res, err := ch.srv.Update(c.Get("user"), commentID, *cnv) // c.Get("user") user <- table "user" yang diambil adalah ID nya
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessResponse(http.StatusOK, "sukses update comment", UpdToResponse(res)))
	}
}

func (ch *commentHandle) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		ParamID := c.Param("comment_id")
		commentID, _ := strconv.Atoi(ParamID)
		err := ch.srv.Delete(c.Get("user"), commentID)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessResponse(http.StatusOK, "sukses menghapus comment", err))
	}
}
