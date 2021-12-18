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

func GetStuApplication(c *gin.Context) {
	username := c.Query("username")
	applications, err := database.GetApplicationsByUserName(username)
	if err != nil {
		logrus.Errorf(constant.Service+"GetStuApplication Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	if applications == nil {
		logrus.Errorf(constant.Service + "GetStuApplication Failed, application is nil")
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	resp := make([]api.StuApplicationResponse, 0)
	for _, a := range applications {
		resp = append(resp, api.StuApplicationResponse{
			ApplicationID: a.ApplicationID,
			UserName:      a.UserName,
			University:    a.University,
			Major:         a.Major,
			IsAdmitted:    a.IsAdmitted,
			IsConfirmed:   a.IsConfirmed,
			SubmitTime:    a.SubmitTime,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":        "OK",
		"code":           0,
		"stuApplication": resp,
	})
}

func PostStuApplication(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	params := make(map[string]interface{})
	c.BindJSON(&params)
	mp := params["stuApplication"].(map[string]interface{})
	applicationID := int64(mp["applicationID"].(float64))
	update := gin.H{
		"university": mp["university"].(string),
		"major":      mp["major"].(string),
	}
	err := database.UpdateApplicationByApplicationID(applicationID, username, update)

	if err != nil {
		logrus.Errorf(constant.Service+"PostStuApplication Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	c.JSON(http.StatusOK, GenResponseWithOK())
}

func DeleteStuApplication(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	params := make(map[string]interface{})
	c.BindJSON(&params)
	applicationID := int64(params["applicationID"].(float64))

	err := database.DeleteApplicationByApplicationID(username, applicationID)

	if err != nil {
		logrus.Errorf(constant.Service+"DeleteStuApplication Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	c.JSON(http.StatusOK, GenResponseWithOK())
}

func PutStuApplication(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	params := make(map[string]interface{})
	c.BindJSON(&params)
	mp := params["stuApplication"].(map[string]interface{})
	uni := mp["university"].(string)
	maj := mp["major"].(string)
	err := database.AddApplication(username, uni, maj)
	if err != nil {
		logrus.Errorf(constant.Service+"PutStuApplication Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	c.JSON(http.StatusOK, GenResponseWithOK())
}

func GetTeaApplication(c *gin.Context){
	uni := c.Query("university")
	maj := c.Query("major")
	applications, err := database.GetApplicationsByUniversityAndMajor(uni,maj)
	if err != nil {
		logrus.Errorf(constant.Service+"GetTeaApplication Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	if applications == nil {
		logrus.Errorf(constant.Service + "GetTeaApplication Failed, application is nil")
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	resp := make([]api.StuApplicationResponse, 0)
	for _, a := range applications {
		resp = append(resp, api.StuApplicationResponse{
			ApplicationID: a.ApplicationID,
			UserName:      a.UserName,
			University:    a.University,
			Major:         a.Major,
			IsAdmitted:    a.IsAdmitted,
			IsConfirmed:   a.IsConfirmed,
			SubmitTime:    a.SubmitTime,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":        "OK",
		"code":           0,
		"teaApplication": resp,
	})
}

func PostTeaAdmit(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	applicationID := int64(params["applicationID"].(float64))

	err := database.AdmitApplicationByApplicationID(applicationID)

	if err != nil {
		logrus.Errorf(constant.Service+"PostTeaAdmit Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	c.JSON(http.StatusOK, GenResponseWithOK())
}

func PostStuConfirm(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	applicationID := int64(params["applicationID"].(float64))

	err := database.ConfirmApplicationByApplicationID(applicationID)

	if err != nil {
		logrus.Errorf(constant.Service+"PostStuConfirm Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	c.JSON(http.StatusOK, GenResponseWithOK())
}