package data

import (
	"errors"
	"log"
	"socialmedia/features/comments"
	"socialmedia/features/comments/data"
	"socialmedia/features/posts"

	"gorm.io/gorm"
)

type postsData struct {
	db *gorm.DB
}

func New(db *gorm.DB) posts.PostData {
	return &postsData{
		db: db,
	}
}

func (pd *postsData) Add(userID int, newPost posts.Core) (posts.Core, error) {
	cnv := CoreToData(newPost)
	cnv.UserID = uint(userID)
	err := pd.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return posts.Core{}, errors.New("querry error,fail to add item")
	}

	newPost.ID = cnv.ID
	return newPost, nil
}

func (pd *postsData) GetPost() ([]posts.Core, error) {
	res := []Posts{}
	err := pd.db.Find(&res).Error
	if err != nil {
		log.Println("no data found")
		return []posts.Core{}, errors.New("data not found")
	}
	result := []posts.Core{}
	for i := 0; i < len(res); i++ {
		temp := res[i]
		result = append(result, ToCore(temp))
		qry := User{}
		err := pd.db.Where("id=?", res[i].UserID).First(&qry).Error
		if err != nil {
			log.Println("no data found")
			return []posts.Core{}, errors.New("data not found")
		}
		result[i].Owner = qry.Username
	}

	return result, nil
}

func (pd *postsData) GetPostDetail(postID int) (interface{}, error) {
	resPost := Posts{}
	qryPost := pd.db.Where("id = ?", postID).First(&resPost)
	err := qryPost.Error
	if err != nil {
		log.Println("no data found")
		return Posts{}, errors.New("data not found")
	}
	resCom := []data.Comments{}
	qryCom := pd.db.Find(&resCom)
	err2 := qryCom.Error
	if err2 != nil {
		log.Println("no data found")
		return []data.Comments{}, errors.New("data not found")
	}

	result := []comments.Core{}
	for i := 0; i < len(resCom); i++ {
		temp := resCom[i]
		result = append(result, ToCoreCom(temp))
		qry := User{}
		err := pd.db.Where("id = ?", resCom[i].UserID).First(&qry).Error
		if err != nil {
			log.Println("no data found")
			return []comments.Core{}, errors.New("data not found")
		}
		result[i].Owner = qry.Username
	}

	return result, nil
}

func (pd *postsData) Update(postID int, userID int, updateData posts.Core) (posts.Core, error) {
	cnv := CoreToData(updateData)
	qry := pd.db.Model(&Posts{}).Where("id = ? AND user_id = ?", postID, userID).Updates(&cnv)

	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return posts.Core{}, errors.New("tidak ada post yang terhapus")
	}

	err := qry.Error
	if err != nil {
		log.Println("update post query error", err.Error())
		return posts.Core{}, err
	}

	return ToCore(cnv), nil
}

func (pd *postsData) Delete(userID int, postID int) error {
	qry := pd.db.Where("id=? AND user_id=?", postID, userID).Delete(&Posts{})
	err := qry.Error

	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return errors.New("tidak ada post yang terhapus")
	}

	if err != nil {
		log.Println("delete query error")
		return errors.New("tidak bisa menghapus data")
	}

	return nil
}
