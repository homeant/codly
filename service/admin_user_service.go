package service

import (
	"errors"
	"github.com/homeanter/codly/config"
	"github.com/homeanter/codly/datastore"
	"github.com/homeanter/codly/model"
	"github.com/homeanter/codly/utils"
)

type AdminUserInterface interface {
	Register(data *model.AdminUserRegister) (adminUser *model.AdminUser, err error)
	GetAdminUser(username string) (data *model.AdminUser, err error)
	Login(data *model.AdminUserLogin) (token *model.AuthToken, err error)
}

type adminUserService struct {
}

func (service *adminUserService) Login(data *model.AdminUserLogin) (token *model.AuthToken, err error) {
	adminUser, err := service.GetAdminUser(data.Username)
	if err != nil {
		return nil, err
	}
	newPassword := utils.EncryptPassword(data.Password, config.Config.BcryptPassword)
	if adminUser.Password != newPassword {
		return nil, errors.New("password error")
	}
	generateToken, err := utils.GenerateToken(adminUser.ID)
	if err != nil {
		return nil, err
	}
	return &model.AuthToken{
		Token:    generateToken,
		Username: adminUser.Username,
	}, nil
}

var AdminUserService AdminUserInterface = &adminUserService{}

func (service *adminUserService) Register(data *model.AdminUserRegister) (adminUser *model.AdminUser, err error) {
	result, err := model.AdminUserDatastore.FetchOne("username=?", data.Username)
	if err != nil {
		return nil, err
	}
	if result != nil {
		return nil, errors.New("username already exists")
	}
	// 对密码进行加密
	newUser := model.AdminUser{
		Username: data.Username,
		Password: utils.EncryptPassword(data.Password, config.Config.BcryptPassword),
		NickName: data.NickName,
	}
	adminUser, err = model.AdminUserDatastore.Create(datastore.Datastore.DB(), &newUser)
	if err != nil {
		return nil, err
	}
	return adminUser, nil
}

func (service *adminUserService) GetAdminUser(username string) (data *model.AdminUser, err error) {
	data, err = model.AdminUserDatastore.FetchOne("username=?", username)
	return
}
