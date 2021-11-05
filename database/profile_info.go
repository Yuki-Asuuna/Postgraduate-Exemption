package database

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/utils/mysql"
	"github.com/sirupsen/logrus"
)

func AddProfile(ID int64, userName string) error {
	profile := ProfileInfo{
		UserName:  userName,
		ProfileID: ID,
	}
	if err := mysql.GetMySQLClient().Create(&profile).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddProfile Failed, err= %v", err)
		return err
	}
	return nil
}
