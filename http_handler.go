package main

import (
	"Postgraduate-Exemption/service"
)

func http_handler_init() {
	r.NoRoute(service.NotFound)
	r.GET("/ping", service.Ping)
	r.POST("/register", service.Register)
	r.POST("/login", service.Login)
}
