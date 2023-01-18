package services

import (
	"errors"
	"log"
	"socialmedia/features/comments"
	"socialmedia/helper"
	"strings"

	"github.com/go-playground/validator"
)

type commentSrv struct {
	data comments.CommentData
	vld  *validator.Validate
}

func New(d comments.CommentData) comments.CommentService {
	return &commentSrv{
		data: d,
		vld:  validator.New(),
	}
}

func (cs *commentSrv) Add(token interface{}, newComment comments.Core) (comments.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return comments.Core{}, errors.New("comment tidak ditemukan")
	}

	err := cs.vld.Struct(newComment)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		return comments.Core{}, errors.New("input comment tidak sesuai dengan arahan")
	}

	res, err := cs.data.Add(userID, newComment)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return comments.Core{}, errors.New(msg)
	}

	return res, nil

}

func (cs *commentSrv) Update(token interface{}, commentID int, updatedData comments.Core) (comments.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return comments.Core{}, errors.New("comment tidak ditemukan")
	}

	res, err := cs.data.Update(userID, commentID, updatedData)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return comments.Core{}, errors.New(msg)
	}

	return res, nil
}

func (cs *commentSrv) Delete(token interface{}, commentID int) error {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return errors.New("data not found")
	}

	err := cs.data.Delete(userID, commentID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return errors.New(msg)
	}

	return nil
}
