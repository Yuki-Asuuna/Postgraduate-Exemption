package service

import (
	"github.com/gin-gonic/gin"
)

func GenResponseWithOK() gin.H {
	return gin.H{
		"message": "OK",
		"code":  0,
	}
}

func GenResponseWithUnauthorized() gin.H {
	return gin.H{
		"message": "UnAuthorized",
		"code":  -2,
	}
}
