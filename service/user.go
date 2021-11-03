package service

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/database"
	"Postgraduate-Exemption/utils/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Register(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	username := params["username"].(string)
	password := params["password"].(string)
	identity := int64(params["identity"].(float64))
	phonenumber := params["phone_number"].(string)
	user, err := database.GetUserByUserName(username)
	if err != nil {
		logrus.Error(constant.Service+"Register Failed, err= %v", err)
		return
	}
	if user != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "UserName already exists",
			"status":  http.StatusConflict,
		})
		return
	}
	if err := database.AddUser(username, password, identity, phonenumber); err != nil {
		logrus.Error(constant.Service+"Register Failed, err= %v", err)
		c.JSON(http.StatusConflict, gin.H{
			"message": "UserName already exists",
			"status":  http.StatusConflict,
		})
		return
	}
	c.JSON(http.StatusOK, GenResponseWithOK())
}

func Login(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	username := params["username"].(string)
	password := params["password"].(string)
	user, err := database.GetUserByUserName(username)
	if err != nil {
		logrus.Error(constant.Service+"Login Failed, err= %v", err)
		return
	}
	if user == nil {
		c.JSON(http.StatusNoContent, gin.H{
			"message": "UserName Not Found",
			"status":  http.StatusNoContent,
		})
		return
	}
	if password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Incorrect Password",
			"status":  http.StatusUnauthorized,
		})
		return
	}
	session, _ := sessions.GetSessionClient().Get(c.Request, "dotcomUser")
	session.Values["authenticated"] = true
	session.Values["username"] = username
	err = sessions.GetSessionClient().Save(c.Request, c.Writer, session)
	if err != nil {
		logrus.Errorf(constant.Service+"Login Failed, err= %v", err)
		return
	}
	c.JSON(http.StatusOK, GenResponseWithOK())
}

func Logout(c *gin.Context) {
	session, _ := sessions.GetSessionClient().Get(c.Request, "dotcomUser")
	session.Values["authenticated"] = false
	err := sessions.GetSessionClient().Save(c.Request, c.Writer, session)
	if err != nil {
		logrus.Errorf(constant.Service+"Logout Failed, err= %v", err)
		return
	}
	c.JSON(http.StatusOK, GenResponseWithOK())
}
