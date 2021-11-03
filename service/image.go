package service

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/utils/helper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	path2 "path"
	"time"
)

const (
	UPLOAD_DIR = "./upload_images"
)

func UploadImage(c *gin.Context) {
	f, err := c.FormFile("image")
	if err != nil {
		logrus.Errorf(constant.Service+"UploadImage Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Upload Image Failed",
			"status":  http.StatusInternalServerError,
		})
		return
	}
	fileName := f.Filename
	ext := path2.Ext(fileName)
	hashName := helper.S2MD5(fileName + time.Now().String())
	path := UPLOAD_DIR + "/" + hashName + ext
	if err := c.SaveUploadedFile(f, path); err != nil {
		logrus.Errorf(constant.Service+"UploadImage Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Upload Image Failed",
			"status":  http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, GenResponseWithOK())
}
