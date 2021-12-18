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
