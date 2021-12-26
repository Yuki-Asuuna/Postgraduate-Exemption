package main

import (
	"Postgraduate-Exemption/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 允许跨域访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
			c.Header("Access-Control-Allow-Origin", "*")  // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func http_handler_init() {
	r.Use(Cors())
	r.NoRoute(service.NotFound)
	r.GET("/ping", service.Ping)
	r.POST("/register", service.Register)
	r.POST("/login", service.Login)
	r.POST("/logout", service.Logout)
	r.POST("/password", service.ChangePassword)
	r.GET("/home", service.AuthMiddleWare(), service.Home)

	// image
	r.POST("/upload_image", service.UploadUserPhoto)
	r.GET("/load_image", service.GetUserPhoto)

	// agreement
	r.GET("/agreement", service.GetAgreementInfo)
	r.POST("/agreement", service.PostAgreementInfo)

	// info
	r.GET("/stu_basic_info", service.GetStudentBasicInfo)
	r.POST("/stu_basic_info", service.PostStudentBasicInfo)
	r.GET("/account_info", service.GetAccountInfo)
	r.GET("/profile_info", service.GetProfileInfo)
	r.POST("/profile_info", service.PostProfileInfo)
	r.GET("/member_info", service.GetMembersInfo)
	r.POST("/member_info", service.PostMemberInfo)
	r.GET("/contact_info", service.GetContactInfo)
	r.POST("/contact_info", service.PostContactInfo)
	r.GET("/study_info", service.GetStudyInfo)
	r.POST("/study_info", service.PostStudyInfo)
	r.GET("/experiences_info", service.GetExperiencesInfo)
	r.POST("/experiences_info", service.PostExperiencesInfo)

	// application
	r.GET("/stu_application", service.GetStuApplication)
	r.DELETE("/stu_application", service.DeleteStuApplication)
	r.POST("/stu_application", service.PostStuApplication)
	r.PUT("/stu_application", service.PutStuApplication)
	r.GET("/tea_application", service.GetTeaApplication)
	r.POST("/tea_admit", service.PostTeaAdmit)
	r.POST("/stu_confirm", service.PostStuConfirm)
}
