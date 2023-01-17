package services

import (
	"errors"
	"socialmedia/features/posts"
	"socialmedia/helper"
	"socialmedia/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	data := mocks.NewPostData(t)
	inputData := posts.Core{
		ID:          uint(0),
		Content:     "ootd",
		Img_content: "",
	}
	resData := posts.Core{
		ID:          uint(1),
		Content:     "ootd",
		Img_content: "",
	}
	t.Run("Success Add", func(t *testing.T) {
		data.On("Add", int(1), inputData).Return(resData, nil).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateJWT(1)
		id := tokenIDUser.(*jwt.Token)
		id.Valid = true
		res, err := srv.Add(id, inputData)
		assert.Nil(t, err)
		assert.Equal(t, inputData.Content, res.Content)
		assert.Equal(t, res.ID, resData.ID)
		data.AssertExpectations(t)
	})
	t.Run("Fail to add item", func(t *testing.T) {
		data.On("Add", int(1), inputData).Return(posts.Core{}, errors.New("internal server error"))
		srv := New(data)
		_, tokenIDUser := helper.GenerateJWT(1)
		id := tokenIDUser.(*jwt.Token)
		id.Valid = true
		res, err := srv.Add(id, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server error")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {

		input := posts.Core{ID: uint(0), Content: "ootd", Img_content: ""}
		data.On("Add", int(1), inputData).Return(posts.Core{}, errors.New("internal server error"))
		srv := New(data)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, input)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("post tidak ditemukan", func(t *testing.T) {
		input := posts.Core{Content: "spongebob", Img_content: "adfasdf"}
		data.On("Add", 1, input).Return(posts.Core{}, errors.New("data not found")).Once()

		srv := New(data)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, input)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	data := mocks.NewPostData(t)
	inputData := posts.Core{
		ID:          uint(0),
		Content:     "ootd",
		Img_content: "",
	}
	resData := posts.Core{
		ID:          uint(1),
		Content:     "ootd2",
		Img_content: "inifoto",
	}
	t.Run("Success Updating", func(t *testing.T) {
		data.On("Update", int(1), int(1), inputData).Return(resData, nil).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateJWT(1)
		UserId := tokenIDUser.(*jwt.Token)
		UserId.Valid = true
		res, err := srv.Update(UserId, 1, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		data.AssertExpectations(t)
	})
	t.Run("Update Fail", func(t *testing.T) {
		data.On("Update", int(1), int(1), inputData).Return(posts.Core{}, errors.New("server error")).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateJWT(1)
		UserId := tokenIDUser.(*jwt.Token)
		UserId.Valid = true
		res, err := srv.Update(UserId, 1, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
	t.Run("Post Not Found", func(t *testing.T) {
		data.On("Update", int(1), int(1), inputData).Return(posts.Core{}, errors.New("posts not found")).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateJWT(1)
		UserId := tokenIDUser.(*jwt.Token)
		UserId.Valid = true
		res, err := srv.Update(UserId, 1, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
	t.Run("jwt tidak valid", func(t *testing.T) {

		input := posts.Core{ID: uint(0), Content: "ootd", Img_content: ""}
		data.On("Add", int(1), input).Return(posts.Core{}, errors.New("internal server error"))
		srv := New(data)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, input)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		assert.Equal(t, uint(0), res.ID)
	})

}

func TestDelete(t *testing.T) {
	repo := mocks.NewPostData(t)

	t.Run("sukses menghapus posts", func(t *testing.T) {
		repo.On("Delete", int(1), int(1)).Return(nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, 1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	// t.Run("jwt tidak valid", func(t *testing.T) {
	// 	// repo.On("Delete", int(-1), int(-1)).Return(posts.Core{}, errors.New("internal server error")).Once()
	// 	_, token := helper.GenerateJWT(1)
	// 	srv := New(repo)

	// 	err := srv.Delete(token, 1)
	// 	assert.NotNil(t, err)
	// 	assert.ErrorContains(t, err, "not found")
	// })

	t.Run("data tidak ditemukan", func(t *testing.T) {
		repo.On("Delete", 5, 1).Return(errors.New("data not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(5)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		repo.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		repo.On("Delete", 1, 1).Return(errors.New("terdapat masalah pada server")).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

}

func TestGetPost(t *testing.T) {
	repo := mocks.NewPostData(t)

	t.Run("Sukses lihat buku", func(t *testing.T) {
		resData := []posts.Core{
			{
				ID:          uint(0),
				Content:     "ootd",
				Img_content: "",
			},
		}

		repo.On("GetPost").Return(resData, nil).Once()

		srv := New(repo)

		res, err := srv.GetPost()
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		repo.On("GetPost").Return([]posts.Core{}, errors.New("data not found")).Once()

		srv := New(repo)
		res, err := srv.GetPost()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		repo.On("GetPost").Return([]posts.Core{}, errors.New("terdapat masalah pada server")).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetPost()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}
