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
		return nil, nil
	}
	return user, nil
}

func AddUser(username string, password string, identity int64, phonenumber string, university string, major string) error {
	user := User{
		UserName:    username,
		Password:    password,
		Identity:    identity,
		PhoneNumber: phonenumber,
		University:  university,
		Major:       major,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	if err := mysql.GetMySQLClient().Create(&user).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddUser Failed, err= %v", err)
		return err
	}
	return nil
}

func UpdateUserByUserName(userName string, updateMap map[string]interface{}) error {
	err := mysql.GetMySQLClient().Model(&User{}).Where("user_name = (?)", userName).Update(updateMap).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdateUserByUserName Failed, err= %v, userName= %v", err, userName)
		return err
	}
	return nil
}
