package database

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/utils/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func AddAgreement(userName string) error {
	agreement := Agreement{
		UserName:        userName,
		HasAgreedNotice: 0,
		HasAgreedHonest: 0,
	}
	if err := mysql.GetMySQLClient().Create(&agreement).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddAgreement Failed, err= %v", err)
		return err
	}
	return nil
}
func GetAgreementByUserName(userName string) (*Agreement, error) {
	agreement := new(Agreement)
	if err := mysql.GetMySQLClient().Where("user_name = ?", userName).Find(&agreement).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Errorf(constant.DAO+"GetAgreementByUserID Failed, err= %v", err)
		return nil, err
	}
	return agreement, nil
}
func UpdateAgreementByUserName(userName string, updateMap map[string]interface{}) error {
	err := mysql.GetMySQLClient().Model(&Agreement{}).Where("user_name = (?)", userName).Update(updateMap).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdateAgreementByUserID Failed, err= %v, userName= %v", err, userName)
		return err
	}
	return nil
}
