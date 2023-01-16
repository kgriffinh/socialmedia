package data

import (
	"errors"
	"log"
	"socialmedia/features/users"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) Login(username string) (users.Core, error) {
	res := User{}

	if err := uq.db.Where("username = ?", username).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return users.Core{}, errors.New("data not found")
	}

	return ToCore(res), nil
}
func (uq *userQuery) Register(newUser users.Core) (users.Core, error) {
	cnv := CoreToData(newUser)
	err := uq.db.Create(&cnv).Error
	if err != nil {
		return users.Core{}, err
	}

	newUser.ID = cnv.ID
	newUser.Password = ""
	return newUser, nil
}
func (uq *userQuery) Profile(id uint) (users.Core, error) {
	res := User{}
	if err := uq.db.Where("id = ?", id).First(&res).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return users.Core{}, err
	}

	return ToCore(res), nil
}

func (uq *userQuery) Update(id uint, updateData users.Core) (users.Core, error) {
	cnv := CoreToData(updateData)
	qry := uq.db.Model(&User{}).Where("id = ?", id).Updates(&cnv)

	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return users.Core{}, errors.New("tidak ada data user yang terhapus")
	}

	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return users.Core{}, err
	}

	return ToCore(cnv), nil
}

func (uq *userQuery) Delete(id uint) error {
	qry := uq.db.Delete(&User{}, id)
	err := qry.Error

	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return errors.New("tidak ada data user yang terhapus")
	}

	if err != nil {
		log.Println("delete query error")
		return errors.New("tidak bisa menghapus data")
	}

	return nil
}
