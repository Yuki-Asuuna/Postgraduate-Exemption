package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var client *gorm.DB

// var server_ip  = "101.34.159.82"
// var port = 3306
// var db_username = "root"
// var db_password = "123456"
// var db_name = "graduate_exemption"
// var charset = "utf8mb4"
func MysqlInit() error {
	var err error
	client, err = gorm.Open("mysql", "root:123456@(101.34.159.82:3306)/graduate_exemption?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	return nil
}

func GetMySQLClient() *gorm.DB {
	return client
}
