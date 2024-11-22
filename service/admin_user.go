package service

import (
	"codly/datastore"
	"codly/model"
)

type AdminUserInterface interface {
	CreateAdminUser(adminUser *model.AdminUser) (data *model.AdminUser, err error)
	GetAdminUser(username string) (data model.AdminUser, err error)
}

type adminUserService struct {
}

var AdminUserService AdminUserInterface = &adminUserService{}

func (service *adminUserService) CreateAdminUser(adminUser *model.AdminUser) (data *model.AdminUser, err error) {
	data, err = model.AdminUserDatastore.Create(datastore.Datastore.DB(), adminUser)
	return
}

func (service *adminUserService) GetAdminUser(username string) (data model.AdminUser, err error) {
	data, err = model.AdminUserDatastore.FetchOne("username=?", username)
	return
}
