package services

import (
	"errors"
	"log"
	"socialmedia/features/posts"
	"socialmedia/helper"
	"strings"
)

type postsSrv struct {
	data posts.PostData
}

func New(d posts.PostData) posts.PostService {
	return &postsSrv{
		data: d,
	}
}

// Add implements book.BookService
func (ps *postsSrv) Add(token interface{}, newPost posts.Core) (posts.Core, error) {
	userID := helper.ExtractToken(token)
	// error token tidak dibutuhkan karena token pasti memiliki user id karena untuk add buku dan sebagainya
	// user perlu login terlebih dahulu sehingga pasti ada id
	// if userID <= 0 {
	// 	return book.Core{}, errors.New("user tidak ditemukan")
	// }
	res, err := ps.data.Add(userID, newPost)
	if err != nil {
		return posts.Core{}, errors.New("something wrong happens,server error")
	}
	return res, nil
}

// Update implements book.BookService
func (ps *postsSrv) Update(token interface{}, postsID int, updatedData posts.Core) (posts.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := ps.data.Update(postsID, userID, updatedData)
	if err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "not found") {
			return posts.Core{}, errors.New("post not found")
		}
		return posts.Core{}, errors.New("internal server error")

	}
	return res, nil
}

// BookList implements book.BookService
func (ps *postsSrv) GetPost() ([]posts.Core, error) {
	res, err := ps.data.GetPost()
	if err != nil {
		log.Println("no result or server error")
		return []posts.Core{}, errors.New("no result or server error")
	}
	// fmt.Println(res)
	return res, nil
}

// Delete implements book.BookService
func (ps *postsSrv) Delete(token interface{}, postsID int) error {
	userID := helper.ExtractToken(token)
	err := ps.data.Delete(userID, postsID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "post") {
			msg = "user dont have any post"
		} else {
			msg = "internal server error"
		}
		return errors.New(msg)
	}
	return nil
}
