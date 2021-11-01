package database

import (
	"Postgraduate-Exemption/utils"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"time"
)

func GetUserByUserName(userName string) (*User, error) {
	user := new(User)
	if err := utils.GetMySQLClient().Where("user_name = ?", userName).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Errorf("[database] GetUserByUserName failed, err= %v", err)
		return nil, err
	}
	return user, nil
}

func AddUser(username string, password string) error {
	user := User{
		UserName:   username,
		Password:   password,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := utils.GetMySQLClient().Create(&user).Error; err != nil {
		logrus.Errorf("[database] AddUser failed, err= %v", err)
		return err
	}
	return nil
}
