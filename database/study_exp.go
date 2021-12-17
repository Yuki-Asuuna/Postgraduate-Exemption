package database

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/utils/mysql"
	"Postgraduate-Exemption/utils/snowflake"
	"github.com/sirupsen/logrus"
)

func AddExperience(studyInfoID int64) error {
	experienceID := snowflake.GenID()
	experience := StudyExp{
		StudyInfoID: studyInfoID,
		ExperienceID: experienceID,
	}
	if err := mysql.GetMySQLClient().Create(&experience).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddExperience Failed, err= %v", err)
		return err
	}
	return nil
}

func GetExperiencesByStudyInfoID(studyInfoID int64) ([]*StudyExp,error) {
	exps := make([]*StudyExp,0)
	err := mysql.GetMySQLClient().Where("study_info_id = (?)", studyInfoID).Find(&exps).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"GetExperiencesByStudyInfoID Failed, err= %v", err)
		return nil, err
	}
	return exps, nil
}

func UpdateExperienceByExperienceID(experienceID int64, updateMap map[string]interface{}) error {
	err := mysql.GetMySQLClient().Model(&StudyExp{}).Where("experience_id = (?)", experienceID).Update(updateMap).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdateExperienceByExperienceID Failed, err= %v, experienceID= %v", err, experienceID)
		return err
	}
	return nil
}