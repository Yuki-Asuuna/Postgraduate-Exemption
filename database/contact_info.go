package database

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/utils/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func AddContactInfo(userName string) error {
	contact := ContactInfo{
		UserName:        userName,
		PhoneNumber:     "0",
		FixedLineNumber: "0",
	}
	if err := mysql.GetMySQLClient().Create(&contact).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddContactInfo Failed, err= %v", err)
		return err
	}
	return nil
}

func GetContactInfoByUserName(userName string) (*ContactInfo, error) {
	contact := new(ContactInfo)
	if err := mysql.GetMySQLClient().Where("user_name = ?", userName).Find(&contact).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Errorf(constant.DAO+"GetContactInfoByUserName Failed, err= %v", err)
		return nil, err
	}
	return contact, nil
}

func UpdateContactInfoByUserName(userName string, updateMap map[string]interface{}) error {
	err := mysql.GetMySQLClient().Model(&ContactInfo{}).Where("user_name = (?)", userName).Update(updateMap).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdateContactInfoByUserName Failed, err= %v, userName= %v", err, userName)
		return err
	}
	return nil
}
