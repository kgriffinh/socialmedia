package services

import (
	"errors"
	"log"

	"socialmedia/features/posts"
	"socialmedia/helper"
	"strings"

	"github.com/go-playground/validator"
)

type postsSrv struct {
	data posts.PostData
	vld  *validator.Validate
}

func New(d posts.PostData) posts.PostService {
	return &postsSrv{
		data: d,
		vld:  validator.New(),
	}
}

func (ps *postsSrv) Add(token interface{}, newPost posts.Core) (posts.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return posts.Core{}, errors.New("user tidak ditemukan")
	}

	err := ps.vld.Struct(newPost)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		return posts.Core{}, errors.New("input content tidak sesuai dengan arahan")
	}

	res, err := ps.data.Add(userID, newPost)
	if err != nil {
		return posts.Core{}, errors.New("something wrong happens,server error")
	}
	return res, nil
}

func (ps *postsSrv) Update(token interface{}, postID int, updatedData posts.Core) (posts.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := ps.data.Update(postID, userID, updatedData)
	if err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "not found") {
			return posts.Core{}, errors.New("content not found")
		}
		return posts.Core{}, errors.New("internal server error")

	}
	return res, nil
}

func (ps *postsSrv) GetPost() ([]posts.Core, error) {
	res, err := ps.data.GetPost()
	if err != nil {
		log.Println("no result or server error")
		return []posts.Core{}, errors.New("no result or server error")
	}

	return res, nil
}

func (ps *postsSrv) GetPostDetail(postID int) (interface{}, error) {
	res, err := ps.data.GetPostDetail(postID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return res, errors.New(msg)
	}

	return res, nil
}

func (ps *postsSrv) Delete(token interface{}, postID int) error {
	userID := helper.ExtractToken(token)
	err := ps.data.Delete(userID, postID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "content") {
			msg = "user dont have any content"
		} else {
			msg = "internal server error"
		}
		return errors.New(msg)
	}
	return nil
}
