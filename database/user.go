package database

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/utils/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"time"
)

func GetUserByUserName(userName string) (*User, error) {
	user := new(User)
	if err := mysql.GetMySQLClient().Where("user_name = ?", userName).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Errorf(constant.DAO+"GetUserByUserName Failed, err= %v", err)
		return nil, err
	}
	return user, nil
}

func AddUser(username string, password string, identity int64, phonenumber string) error {
	user := User{
		UserName:    username,
		Password:    password,
		Identity:    identity,
		PhoneNumber: phonenumber,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	if err := mysql.GetMySQLClient().Create(&user).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddUser Failed, err= %v", err)
		return err
	}
	return nil
}
