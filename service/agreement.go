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

func GetAgreementInfo(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	agreement, err := database.GetAgreementByUserName(username)
	if err != nil {
		logrus.Errorf(constant.Service+"GetAgreementInfo Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	if agreement == nil {
		logrus.Errorf(constant.Service + "GetAgreementInfo Failed, profile is nil")
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	resp := api.GetAgreementResponse{
		HasAgreedHonest: agreement.HasAgreedHonest,
		HasAgreedNotice: agreement.HasAgreedNotice,
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "OK",
		"code":      0,
		"agreement": resp,
	})
}

func PostAgreementInfo(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	params := make(map[string]interface{})
	c.BindJSON(&params)
	//fmt.Println(params["stuBasicInfo"])
	mp := params["agreement"].(map[string]interface{})
	update := gin.H{
		"has_agreed_notice": int64(mp["hasAgreedNotice"].(float64)),
		"has_agreed_honest": int64(mp["hasAgreedHonest"].(float64)),
	}
	err := database.UpdateAgreementByUserName(username, update)

	if err != nil {
		logrus.Errorf(constant.Service+"PostAgreementInfo Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	c.JSON(http.StatusOK, GenResponseWithOK())
}
