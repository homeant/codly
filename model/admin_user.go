package model

import (
	"codly/datastore"
	"gorm.io/gorm"
)

type AdminUser struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;comment:主键ID" json:"id"`
	Username string `gorm:"unique;comment:用户名" json:"username"`
	Password string `gorm:"comment:密码" json:"password"`
}

type AdminUserInterface interface {
	Create(tx *gorm.DB, data *AdminUser) (user *AdminUser, err error)
	Update(tx *gorm.DB, id uint, data map[string]interface{}) (rowsAffected int64, err error)
	Delete(tx *gorm.DB, data []int) (rowsAffected int64, err error)
	FetchOne(query interface{}, args ...interface{}) (record AdminUser, err error)
}

type store struct {
	tx *gorm.DB
}

func (u store) Create(tx *gorm.DB, data *AdminUser) (user *AdminUser, err error) {
	db := tx.Create(data)
	if err = db.Error; db.Error != nil {
		return
	}
	return data, err
}

func (u store) Update(tx *gorm.DB, id uint, data map[string]interface{}) (rowsAffected int64, err error) {
	db := tx.Model(&AdminUser{}).Where("id = ?", id).Updates(data)
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (u store) Delete(tx *gorm.DB, data []int) (rowsAffected int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (u store) FetchOne(query interface{}, args ...interface{}) (record AdminUser, err error) {
	db := u.tx.Where(query, args...).First(&record)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

var AdminUserDatastore AdminUserInterface = &store{tx: datastore.Datastore.DB()}
