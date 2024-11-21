package service

import (
	"codly/model"
	"gorm.io/gorm"
)

type AdminUserService struct {
	db *gorm.DB
}

func NewAdminUserService(db *gorm.DB) *AdminUserService {
	return &AdminUserService{db: db}
}

func (service *AdminUserService) CreateAdminUser(adminUser model.AdminUser) {
	service.db.Create(&adminUser)
}
