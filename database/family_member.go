package database

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/utils/mysql"
	"github.com/sirupsen/logrus"
)

func AddFamilyMember(profileID int64) error {
	member := FamilyMember{
		ProfileID: profileID,
	}
	if err := mysql.GetMySQLClient().Create(&member).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddFamilyNumber Failed, err= %v", err)
		return err
	}
	return nil
}

func GetFamilyMembersByProfileID(profileID int64) ([]*FamilyMember, error) {
	members := make([]*FamilyMember, 0)
	err := mysql.GetMySQLClient().Where("profile_id = (?)", profileID).Find(&members).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"GetFamilyMembersByProfileID Failed, err= %v", err)
		return nil, err
	}
	return members, nil
}

func UpdateFamilyMemberByProfileID(profileID int64, updateMap map[string]interface{}) error {
	err := mysql.GetMySQLClient().Model(&FamilyMember{}).Where("profile_id = (?)", profileID).Update(updateMap).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdateFamilyMemberByProfileID Failed, err= %v, profileID= %v", err, profileID)
		return err
	}
	return nil
}
