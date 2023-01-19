package services

import (
	"errors"
	"log"
	"mime/multipart"

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

func (ps *postsSrv) Add(token interface{}, file multipart.FileHeader, newPost posts.Core) (posts.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return posts.Core{}, errors.New("user tidak ditemukan")
	}

	formFile, err := file.Open()
	if err != nil {
		return posts.Core{}, errors.New("input tidak sesuai")
	}

	uploadUrl, err := helper.NewMediaUpload().FileUpload(helper.File{File: formFile})
	if err != nil {
		return posts.Core{}, errors.New("input tidak sesuai")
	}

	newPost.Img_content = uploadUrl

	err2 := ps.vld.Struct(newPost)
	if err2 != nil {
		if _, ok := err2.(*validator.InvalidValidationError); ok {
			log.Println(err2)
		}
		return posts.Core{}, errors.New("input content tidak sesuai dengan arahan")
	}

	res, err2 := ps.data.Add(userID, newPost)
	if err2 != nil {
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
		return nil, errors.New(msg)
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
