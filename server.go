package main

import (
	"Postgraduate-Exemption/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var r = gin.Default()

func main() {
	// init api handler
	http_handler_init()

	// init database
	if err := utils.MysqlInit(); err != nil {
		logrus.Error(err)
	}

	// start gin
	if err := r.Run(":8000"); err != nil {
		logrus.Error(err)
	}
}
