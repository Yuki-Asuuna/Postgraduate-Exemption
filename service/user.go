package service

import (
	"Postgraduate-Exemption/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Register(c *gin.Context) {
	//password := c.PostForm("password")
	//username := c.PostForm("username")
	params := make(map[string]interface{})
	c.BindJSON(&params)
	username := params["username"].(string)
	password := params["password"].(string)
	user, err := database.GetUserByUserName(username)
	if err != nil {
		logrus.Error("[service] Register failed, err= %v", err)
		return
	}
	if user != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "UserName already exists",
			"status":  http.StatusConflict,
		})
		return
	}
	if err := database.AddUser(username, password); err != nil {
		logrus.Error("[service] Register failed, err= %v", err)
		c.JSON(http.StatusConflict, gin.H{
			"message": "UserName already exists",
			"status":  http.StatusConflict,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"status":  http.StatusOK,
	})
}

func Login(c *gin.Context) {

}