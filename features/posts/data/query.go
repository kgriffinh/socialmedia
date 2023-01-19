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
		result[i].Username = qry.Username
		result[i].Userpp = qry.Userpp
	}
	return result, nil
}

func (pd *postsData) GetPostDetail(postID int) (interface{}, error) {
	res := Posts{}
	err := pd.db.Preload("Comment").Where("id = ?", postID).Find(&res).Error
	if err != nil {
		log.Println("no data found")
		return nil, errors.New("data not found")
	}

	resUser := User{}
	if err := pd.db.Where("id = ?", res.UserID).First(&resUser).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return nil, err
	}

	result := posts.Core{
		ID:          res.ID,
		Content:     res.Content,
		Img_content: res.Img_content,
		Username:    resUser.Username,
		Userpp:      resUser.Userpp,
		Comment:     res.Comment,
	}

	return result, nil
	// qryPost := pd.db.Where("id = ?", postID).Preload("Comments").First(&resPost)
	// res := Posts{}
	// err := pd.db.Preload("Comment").Where("id = ?", postID).Find(&res).Error

	// if err != nil {
	// 	log.Println("no data found")
	// 	return posts.Core{}, errors.New("data not found")
	// }

	// return posts.Core{}, nil
	// resPost := map[string]interface{}{}
	// qryPost := pd.db.Model(&Posts{}).Raw("SELECT p.id, p.content, p.img_content, u.username FROM posts p JOIN users u ON u.id = p.user_id").Scan(&resPost)
	// // Select("content").Where("id = ?", postID).First(&resPost)
	// err := qryPost.Error
	// if err != nil {
	// 	log.Println("no data found")
	// 	return nil, errors.New("data not found")
	// }
	// resCom := map[string]interface{}{}
	// qryCom := pd.db.Model(&Comment{}).Raw("SELECT comments.id, comments.text, users.username FROM comments JOIN users ON users.id = comments.user_id WHERE comments.post_id = ?", postID).Scan(&resCom)
	// err2 := qryCom.Error
	// if err2 != nil {
	// 	log.Println("no data found")
	// 	return nil, errors.New("data not found")
	// }

	// result := []comments.Core{}
	// for i := 0; i < len(resCom); i++ {
	// 	temp := resCom[i]
	// 	result = append(result, ToCoreCom(temp))
	// 	qry := User{}
	// 	err := pd.db.Where("id = ?", resCom[i].UserID).First(&qry).Error
	// 	if err != nil {
	// 		log.Println("no data found")
	// 		return []comments.Core{}, errors.New("data not found")
	// 	}
	// 	result[i].Owner = qry.Username
	// }
	// resPost["comments"] = resCom

	// return resPost, nil
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
