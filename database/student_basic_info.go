package database

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/utils/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func AddStudentBasicInfo(userName string) error {
	stu := StudentBasicInfo{
		UserName: userName,
	}
	if err := mysql.GetMySQLClient().Create(&stu).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddStudentBasicInfo Failed, err= %v", err)
		return err
	}
	return nil
}

func GetStudentBasicInfoByUserName(userName string) (*StudentBasicInfo, error) {
	stu := new(StudentBasicInfo)
	if err := mysql.GetMySQLClient().Where("user_name = ?", userName).Find(&stu).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Errorf(constant.DAO+"GetStudentBasicInfoByUserName Failed, err= %v", err)
		return nil, err
	}
	return stu, nil
}

func UpdateStudentBasicInfoByUserName(userName string, updateMap map[string]interface{}) error {
	err := mysql.GetMySQLClient().Model(&StudentBasicInfo{}).Where("user_name = (?)", userName).Update(updateMap).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdateStudentBasicInfoByUserName Failed, err= %v, userName= %v", err, userName)
		return err
	}
	return nil
}
