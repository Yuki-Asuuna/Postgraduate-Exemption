package database

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/utils/mysql"
	"Postgraduate-Exemption/utils/snowflake"
	"github.com/sirupsen/logrus"
	"time"
)

func AddApplication(userName string, university string, major string) error {
	applicationID := snowflake.GenID()
	application := Application{
		UserName:      userName,
		ApplicationID: applicationID,
		University:    university,
		Major:         major,
		IsAdmitted:    0,
		isConfirmed:   0,
		SubmitTime:    time.Now(),
	}
	if err := mysql.GetMySQLClient().Create(&application).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddApplication Failed, err= %v", err)
		return err
	}
	return nil
}

func GetApplicationsByUserName(userName string) ([]*Application,error){
	applications := make([]*Application, 0)
	err := mysql.GetMySQLClient().Where("user_name = (?)", userName).Find(&applications).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"GetApplicationsByUserName Failed, err= %v", err)
		return nil, err
	}
	return applications, nil
}

func UpdateApplicationByApplicationID(applicationID int64, updateMap map[string]interface{}) error {
	err := mysql.GetMySQLClient().Model(&Application{}).Where("application_id = (?)", applicationID).Update(updateMap).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdateApplicationByApplicationID Failed, err= %v, applicationID= %v", err, applicationID)
		return err
	}
	return nil
}

func DeleteApplicationByApplicationID(applicationID int64) error {
	err := mysql.GetMySQLClient().Where("application_id = (?)",applicationID).Delete(&Application{}).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"DeleteApplicationByApplicationID Failed, err= %v, applicationID= %v", err, applicationID)
		return err
	}
	return nil
}