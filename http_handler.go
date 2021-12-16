package main

import (
	"Postgraduate-Exemption/service"
)

func http_handler_init() {
	r.NoRoute(service.NotFound)
	r.GET("/ping", service.Ping)
	r.POST("/register", service.Register)
	r.POST("/login", service.Login)
	r.POST("/logout", service.Logout)
	r.GET("/home", service.AuthMiddleWare(), service.Home)

	// image
	r.POST("/upload_image", service.UploadImageTest)

	// info
	r.GET("/stu_basic_info",service.StudentBasicInfo)
	r.GET("/account_info",service.GetAccountInfo)
}
