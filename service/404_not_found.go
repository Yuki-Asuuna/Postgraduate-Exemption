package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFound(c *gin.Context){
	c.JSON(http.StatusNotFound,gin.H{
		"message":"404 page not found",
		"status":404,
	})
}