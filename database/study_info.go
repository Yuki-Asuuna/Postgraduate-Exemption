package database

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/utils/mysql"
	"Postgraduate-Exemption/utils/snowflake"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func AddStudyInfo(userName string) error {
	studyInfoID := snowflake.GenID()
	studyInfo := StudyInfo{
		UserName:    userName,
		StudyInfoID: studyInfoID,
	}
	if err := mysql.GetMySQLClient().Create(&studyInfo).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddStudyInfo Failed, err= %v", err)
		return err
	}
	for i := 0; i < 4; i++ {
		err := AddExperience(studyInfoID)
		if err != nil {
			logrus.Errorf(constant.DAO+"AddStudyInfo Failed, err= %v", err)
			return err
		}
	}
	return nil
}

func GetStudyInfoByUserName(userName string) (*StudyInfo, error) {
	studyInfo := new(StudyInfo)
	if err := mysql.GetMySQLClient().Where("user_name = ?", userName).Find(&studyInfo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Errorf(constant.DAO+"GetStudyInfoByUserName Failed, err= %v", err)
		return nil, err
	}
	return studyInfo, nil
}

func UpdateStudyInfoByUserName(userName string, updateMap map[string]interface{}) error {
	err := mysql.GetMySQLClient().Model(&StudyInfo{}).Where("user_name = (?)", userName).Update(updateMap).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdateStudyInfoByUserName Failed, err= %v, userName= %v", err, userName)
		return err
	}
	return nil
}
