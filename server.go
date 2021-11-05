package main

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/utils/mysql"
	"Postgraduate-Exemption/utils/sessions"
	"Postgraduate-Exemption/utils/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var r = gin.Default()

func logger_init() error {
	logrus.SetReportCaller(true)
	file_writer, err := os.OpenFile("./log/sys.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic("Cannot Open sys.log")
		return err
	}
	logrus.SetOutput(io.MultiWriter(os.Stdout, file_writer))
	return nil
}

func main() {
	// init logger
	logger_init()

	// init api handler
	http_handler_init()
	// init unique id generator (twitter/snowflake)
	if err := snowflake.SnowflakeInit(); err != nil {
		logrus.Errorf(constant.Main+"Init Snowflake Failed, err= %v", err)
		return
	}

	// init session
	if err := sessions.SessionInit(); err != nil {
		logrus.Errorf(constant.Main+"Init Session Failed, err= %v", err)
		return
	}
	logrus.Infof(constant.Main + "Init Session Success!")

	// init mysql database
	if err := mysql.MysqlInit(); err != nil {
		logrus.Error(constant.Main+"Init Mysql Failed, err= %v", err)
		return
	}
	logrus.Infof(constant.Main + "Init Mysql Success!")

	// start gin
	if err := r.Run(":8000"); err != nil {
		logrus.Error(constant.Main+"Run Gin Server Failed, err= %v", err)
	}
}
