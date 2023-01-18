package data

import (
	"errors"
	"log"
	"socialmedia/features/comments"

	"gorm.io/gorm"
)

type commentData struct {
	db *gorm.DB
}

func New(db *gorm.DB) comments.CommentData {
	return &commentData{
		db: db,
	}
}

func (cd *commentData) Add(userID int, newComment comments.Core) (comments.Core, error) {
	cnv := CoreToData(newComment)
	cnv.UserID = uint(userID)
	err := cd.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return comments.Core{}, errors.New("querry error, fail to add item")
	}

	newComment.ID = cnv.ID
	return newComment, nil
}

func (cd *commentData) Update(commentID int, postID int, updateData comments.Core) (comments.Core, error) {
	cnv := CoreToData(updateData)
	qry := cd.db.Model(&Comments{}).Where("id = ? AND post_id = ?", commentID, postID).Updates(&cnv)

	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return comments.Core{}, errors.New("tidak ada comment yang terhapus")
	}

	err := qry.Error
	if err != nil {
		log.Println("update comment query error", err.Error())
		return comments.Core{}, err
	}

	return ToCore(cnv), nil
}

func (cd *commentData) Delete(commentID int, postID int) error {
	qry := cd.db.Where("id = ? AND post_id = ?", commentID, postID).Delete(&Comments{})
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
