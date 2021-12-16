package service

import (
	"Postgraduate-Exemption/api"
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/database"
	"Postgraduate-Exemption/utils/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func StudentBasicInfo(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	info, err := database.GetStudentBasicInfoByUserName(username)
	if err != nil {
		logrus.Error(constant.Service+"Get StudentBasicInfo Failed, err= %v", err)
		return
	}
	if info == nil {
		logrus.Error(constant.Service + "StudentBasicInfo is nil")
		return
	}
	resp := api.GetStudentBasicInfoResponse{
		UserName:               username,
		IdentityNumber:         info.IdentityNumber,
		Name:                   info.Name,
		NamePinyin:             info.NamePinyin,
		MilitaryType:           info.MilitaryType,
		PoliticalStatus:        info.PoliticalStatus,
		Gender:                 info.Gender,
		MartialStatus:          info.MartialStatus,
		BirthLocationProvince:  info.BirthLocationProvince,
		BirthLocationCity:      info.BirthLocationCity,
		BirthLocationCounty:    info.BirthLocationCounty,
		NativeLocationProvince: info.NativeLocationProvince,
		NativeLocationCity:     info.NativeLocationCity,
		NativeLocationCounty:   info.NativeLocationCounty,
	}
	c.JSON(http.StatusOK, gin.H{
		"StuBasicInfo": resp,
		"message":      "OK",
		"code":         0,
	})
}
