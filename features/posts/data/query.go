package data

import (
	"errors"
	"log"
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
	// log.Println(result)
	return result, nil
}

func (pd *postsData) Update(postID int, updateData posts.Core) (posts.Core, error) {
	cnv := CoreToData(updateData)
	qry := pd.db.Model(&User{}).Where("id = ?", postID).Updates(&cnv)

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
