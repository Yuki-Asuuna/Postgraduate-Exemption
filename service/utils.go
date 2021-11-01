package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenResponseWithOK() gin.H {
	return gin.H{
		"message": "OK",
		"status":  http.StatusOK,
	}
}

func GenResponseWithUnauthorized() gin.H {
	return gin.H{
		"message": "UnAuthorized",
		"status":  http.StatusUnauthorized,
	}
}
