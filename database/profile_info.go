package database

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/utils/mysql"
	"Postgraduate-Exemption/utils/snowflake"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func AddProfile(userName string) error {
	ID := snowflake.GenID()
	profile := ProfileInfo{
		UserName:  userName,
		ProfileID: ID,
	}
	if err := mysql.GetMySQLClient().Create(&profile).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddProfile Failed, err= %v", err)
		return err
	}
	for i := 0; i < 3; i++ {
		err := AddFamilyMember(ID)
		if err != nil {
			logrus.Errorf(constant.DAO+"AddProfile Failed, err= %v", err)
			return err
		}
	}
	return nil
}

func GetProfileByUserName(userName string) (*ProfileInfo, error) {
	profile := new(ProfileInfo)
	if err := mysql.GetMySQLClient().Where("user_name = ?", userName).Find(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Errorf(constant.DAO+"GetProfileByUserName Failed, err= %v", err)
		return nil, err
	}
	return profile, nil
}

func UpdateProfileInfoByUserName(userName string, updateMap map[string]interface{}) error {
	err := mysql.GetMySQLClient().Model(&ProfileInfo{}).Where("user_name = (?)", userName).Update(updateMap).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdateProfileInfoByUserName Failed, err= %v, userName= %v", err, userName)
		return err
	}
	return nil
}
