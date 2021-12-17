package database

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/utils/mysql"
	"Postgraduate-Exemption/utils/snowflake"
	"github.com/sirupsen/logrus"
)

func AddFamilyMember(profileID int64) error {
	memberID := snowflake.GenID()
	member := FamilyMember{
		ProfileID: profileID,
		MemberID: memberID,
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

func UpdateFamilyMemberByMemberID(memberID int64, updateMap map[string]interface{}) error {
	err := mysql.GetMySQLClient().Model(&FamilyMember{}).Where("member_id = (?)", memberID).Update(updateMap).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdateFamilyMemberByMemberID Failed, err= %v, memberID= %v", err, memberID)
		return err
	}
	return nil
}
