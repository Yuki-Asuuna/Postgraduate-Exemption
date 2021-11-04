package service

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/database"
	"Postgraduate-Exemption/utils/helper"
	"Postgraduate-Exemption/utils/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"net/http"
	path2 "path"
	"time"
)

const (
	UPLOAD_DIR = "./upload_images"
)

func genInternalServerErrorResponse() gin.H {
	return gin.H{
		"message": "Upload Image Failed",
		"status":  http.StatusInternalServerError,
	}
}

func UploadImage(c *gin.Context, file *multipart.FileHeader) (int64, error) {
	fileName := file.Filename
	ext := path2.Ext(fileName)
	hashName := helper.S2MD5(fileName + time.Now().String())
	path := UPLOAD_DIR + "/" + hashName + ext
	image_id := snowflake.GenID()
	if err := database.AddImage(image_id, path); err != nil {
		logrus.Errorf(constant.Service+"UploadImage Failed, err= %v", err)
		return 0, nil
	}
	if err := c.SaveUploadedFile(file, path); err != nil {
		logrus.Errorf(constant.Service+"UploadImage Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, genInternalServerErrorResponse())
	}
	return image_id, nil
}

func UploadImageTest(c *gin.Context) {
	f, err := c.FormFile("image")
	if err != nil {
		logrus.Errorf(constant.Service+"UploadImage Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, genInternalServerErrorResponse())
		return
	}
	ID, err := UploadImage(c, f)
	if err != nil {
		logrus.Errorf(constant.Service+"UploadImage Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, genInternalServerErrorResponse())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"pic_id":  ID,
		"status":  http.StatusOK,
	})
}
