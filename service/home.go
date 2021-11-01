package service

import (
	"Postgraduate-Exemption/utils/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello," + sessions.GetUserNameBySession(c),
		"status":  http.StatusOK,
	})
}
