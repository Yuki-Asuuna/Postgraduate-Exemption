package service

import (
	"Postgraduate-Exemption/api"
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/database"
	"Postgraduate-Exemption/utils/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetStudentBasicInfo(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	info, err := database.GetStudentBasicInfoByUserName(username)
	if err != nil {
		logrus.Error(constant.Service+"GetStudentBasicInfo Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	if info == nil {
		logrus.Error(constant.Service + "StudentBasicInfo is nil")
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	resp := api.GetStudentBasicInfoResponse{
		UserName:               username,
		IdentityNumber:         info.IdentityNumber,
		Name:                   info.Name,
		NamePinyin:             info.NamePinyin,
		MilitaryType:           info.MilitaryType,
		PoliticalStatus:        info.PoliticalStatus,
		Gender:                 info.Gender,
		MartialStatus:          info.MartialStatus,
		BirthLocationProvince:  info.BirthLocationProvince,
		BirthLocationCity:      info.BirthLocationCity,
		BirthLocationCounty:    info.BirthLocationCounty,
		NativeLocationProvince: info.NativeLocationProvince,
		NativeLocationCity:     info.NativeLocationCity,
		NativeLocationCounty:   info.NativeLocationCounty,
		ImageID:                info.ImageID,
	}
	c.JSON(http.StatusOK, gin.H{
		"stuBasicInfo": resp,
		"message":      "OK",
		"code":         0,
	})
}

func PostStudentBasicInfo(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	params := make(map[string]interface{})
	c.BindJSON(&params)
	//fmt.Println(params["stuBasicInfo"])
	mp := params["stuBasicInfo"].(map[string]interface{})
	update := gin.H{
		"identity_number":          mp["identityNumber"].(string),
		"name":                     mp["name"].(string),
		"name_pinyin":              mp["namePinYin"].(string),
		"military_type":            int64(mp["militaryType"].(float64)),
		"political_status":         int64(mp["politicalStatus"].(float64)),
		"gender":                   int64(mp["gender"].(float64)),
		"martial_status":           int64(mp["martialStatus"].(float64)),
		"birth_location_province":  mp["birthLocationProvince"].(string),
		"birth_location_city":      mp["birthLocationCity"].(string),
		"birth_location_county":    mp["birthLocationCounty"].(string),
		"native_location_province": mp["nativeLocationProvince"].(string),
		"native_location_city":     mp["nativeLocationCity"].(string),
		"native_location_county":   mp["nativeLocationCounty"].(string),
		"image_id":                 int64(mp["imageID"].(float64)),
	}
	err := database.UpdateStudentBasicInfoByUserName(username, update)

	if err != nil {
		logrus.Errorf(constant.Service+"PostStudentBasicInfo Failed, err= %v", err)
		return
	}
	c.JSON(http.StatusOK, GenResponseWithOK())
}

func GetAccountInfo(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	account, err := database.GetUserByUserName(username)
	if err != nil {
		logrus.Errorf(constant.Service+"GetAccountInfo Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	resp := api.GetAccountInfoResponse{
		UserName:    account.UserName,
		PhoneNumber: account.PhoneNumber,
		Identity:    account.Identity,
		University:  account.University,
		Major:       account.Major,
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "OK",
		"code":        0,
		"accountInfo": resp,
	})
}

func GetProfileInfo(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	profile, err := database.GetProfileByUserName(username)
	if err != nil {
		logrus.Errorf(constant.Service+"GetProfileInfo Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	resp := api.GetProfileInfoResponse{
		UserName:                  profile.UserName,
		ProfileID:                 profile.ProfileID,
		ProfileLocationProvince:   profile.ProfileLocationProvince,
		ProfileLocationCity:       profile.ProfileLocationCity,
		ProfileLocationCounty:     profile.ProfileLocationCounty,
		ProfileAddress:            profile.ProfileAddress,
		PostCode:                  profile.PostCode,
		ResidenceLocationProvince: profile.ResidenceLocationProvince,
		ResidenceLocationCity:     profile.ResidenceLocationCity,
		ResidenceLocationCounty:   profile.ResidenceLocationCounty,
		ResidenceAddress:          profile.ResidenceAddress,
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "OK",
		"code":        0,
		"profileInfo": resp,
	})
}

func PostProfileInfo(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	params := make(map[string]interface{})
	c.BindJSON(&params)
	//fmt.Println(params["stuBasicInfo"])
	mp := params["profileInfo"].(map[string]interface{})
	update := gin.H{
		"profile_location_province":   mp["profileLocationProvince"].(string),
		"profile_location_city":       mp["profileLocationCity"].(string),
		"profile_location_county":     mp["profileLocationCounty"].(string),
		"profile_address":             mp["profileAddress"].(string),
		"post_code":                   mp["postCode"].(string),
		"residence_location_province": mp["residenceLocationProvince"].(string),
		"residence_location_city":     mp["residenceLocationCity"].(string),
		"residence_location_county":   mp["residenceLocationCounty"].(string),
		"residence_address":           mp["residenceAddress"].(string),
	}
	err := database.UpdateProfileInfoByUserName(username, update)

	if err != nil {
		logrus.Errorf(constant.Service+"PostProfileInfo Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	c.JSON(http.StatusOK, GenResponseWithOK())
}

func GetMembersInfo(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	profile, err := database.GetProfileByUserName(username)
	if err != nil {
		logrus.Errorf(constant.Service+"GetMemberInfo Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	if profile == nil {
		logrus.Errorf(constant.Service + "GetMemberInfo Failed, profile is nil")
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	members, err := database.GetFamilyMembersByProfileID(profile.ProfileID)
	resp := make([]api.GetMemberInfoResponse, 0)
	for _, m := range members {
		resp = append(resp, api.GetMemberInfoResponse{
			MemberID:     m.MemberID,
			ProfileID:    m.ProfileID,
			Name:         m.Name,
			Relationship: m.Relationship,
			PhoneNumber:  m.PhoneNumber,
			Address:      m.Address,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "OK",
		"code":        0,
		"membersInfo": resp,
	})
}

func PostMemberInfo(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	//fmt.Println(params["stuBasicInfo"])
	mp := params["memberInfo"].(map[string]interface{})
	update := gin.H{
		"name":         mp["name"].(string),
		"relationship": mp["relationship"].(string),
		"phone_number": mp["phoneNumber"].(string),
		"address":      mp["address"].(string),
	}
	err := database.UpdateFamilyMemberByMemberID(int64(mp["memberID"].(float64)), update)

	if err != nil {
		logrus.Errorf(constant.Service+"PostMemberInfo Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	c.JSON(http.StatusOK, GenResponseWithOK())
}

func GetContactInfo(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	contact, err := database.GetContactInfoByUserName(username)
	if err != nil {
		logrus.Errorf(constant.Service+"GetContactInfo Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	if contact == nil {
		logrus.Errorf(constant.Service + "GetContactInfo Failed, profile is nil")
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	resp := api.GetContactInfoResponse{
		PhoneNumber:     contact.PhoneNumber,
		FixedLineNumber: contact.FixedLineNumber,
		Address:         contact.Address,
		PostCode:        contact.PostCode,
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "OK",
		"code":        0,
		"contactInfo": resp,
	})
}

func PostContactInfo(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	params := make(map[string]interface{})
	c.BindJSON(&params)
	//fmt.Println(params["stuBasicInfo"])
	mp := params["contactInfo"].(map[string]interface{})
	update := gin.H{
		"phone_number":      mp["phoneNumber"].(string),
		"fixed_line_number": mp["fixedLineNumber"].(string),
		"address":           mp["address"].(string),
		"post_code":         mp["postCode"].(string),
	}
	err := database.UpdateContactInfoByUserName(username, update)

	if err != nil {
		logrus.Errorf(constant.Service+"PostContactInfo Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	c.JSON(http.StatusOK, GenResponseWithOK())
}

func GetStudyInfo(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	study, err := database.GetStudyInfoByUserName(username)
	if err != nil {
		logrus.Errorf(constant.Service+"GetStudyInfo Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	if study == nil {
		logrus.Errorf(constant.Service + "GetStudyInfo Failed, studyInfo is nil")
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	resp := api.GetStudyInfoResponse{
		StudyInfoID:     study.StudyInfoID,
		UserName:        study.UserName,
		SchoolName:      study.SchoolName,
		Writing:         study.Writing,
		AwardPunishment: study.AwardPunishment,
		Cheating:        study.Cheating,
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "OK",
		"code":      0,
		"studyInfo": resp,
	})
}

func PostStudyInfo(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	params := make(map[string]interface{})
	c.BindJSON(&params)
	mp := params["studyInfo"].(map[string]interface{})
	update := gin.H{
		"school_name":      mp["schoolName"].(string),
		"writing":          mp["writing"].(string),
		"award_punishment": mp["awardPunishment"].(string),
		"cheating":         mp["cheating"].(string),
	}
	err := database.UpdateStudyInfoByUserName(username, update)

	if err != nil {
		logrus.Errorf(constant.Service+"PostStudyInfo Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	c.JSON(http.StatusOK, GenResponseWithOK())
}

func GetExperiencesInfo(c *gin.Context) {
	username := sessions.GetUserNameBySession(c)
	studyInfo, err := database.GetStudyInfoByUserName(username)
	if err != nil {
		logrus.Errorf(constant.Service+"GetExperienceInfo Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	if studyInfo == nil {
		logrus.Errorf(constant.Service + "GetExperienceInfo Failed, profile is nil")
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	exps, err := database.GetExperiencesByStudyInfoID(studyInfo.StudyInfoID)
	resp := make([]api.GetExperiencesInfoResponse, 0)
	for _, e := range exps {
		resp = append(resp, api.GetExperiencesInfoResponse{
			ExperienceID: e.ExperienceID,
			Interval:     e.Interval,
			WorkPlace:    e.WorkPlace,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":        "OK",
		"code":           0,
		"experiencesInfo": resp,
	})
}

func PostExperiencesInfo(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	//fmt.Println(params["stuBasicInfo"])
	mp := params["experiencesInfo"].(map[string]interface{})
	update := gin.H{
		"interval":   mp["interval"].(string),
		"work_place": mp["workPlace"].(string),
	}
	err := database.UpdateExperienceByExperienceID(int64(mp["experienceID"].(float64)), update)

	if err != nil {
		logrus.Errorf(constant.Service+"PostExperiencesInfo Failed, err= %v", err)
		c.JSON(http.StatusInternalServerError, GenResponseWithDatabaseFailed())
		return
	}
	c.JSON(http.StatusOK, GenResponseWithOK())
}
