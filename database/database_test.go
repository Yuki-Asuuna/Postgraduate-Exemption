package database

import (
	"Postgraduate-Exemption/utils/mysql"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	mysql.MysqlInit()
	m.Run()
	os.Exit(0)
}

func TestGetImage(t *testing.T) {
	//t.Skip()
	ret, err := GetImages([]int64{244115596130648064, 244116453656100864, 0})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ret)
}

func TestAddAgreement(t *testing.T) {
	err := AddAgreement("edward")
	if err != nil {
		t.Error(err)
	}
}

func TestGetAgreementByUserName(t *testing.T) {
	agreement, err := GetAgreementByUserName("edward")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(agreement)
}

func TestUpdateAgreementByUserName(t *testing.T) {
	err := UpdateAgreementByUserName("edward", map[string]interface{}{
		"has_agreed_notice": 1,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestAddContactInfo(t *testing.T) {
	err := AddContactInfo("jack")
	if err != nil {
		t.Error(err)
	}
}

func TestGetAddContactInfo(t *testing.T) {
	contact, err := GetContactInfoByUserName("jack")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(contact)
}

func TestUpdateContactInfoByUserName(t *testing.T) {
	err := UpdateContactInfoByUserName("jack", map[string]interface{}{
		"phone_number": "10086",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestAddStudentBasicInfo(t *testing.T) {
	err := AddStudentBasicInfo("jack")
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateStudentBasicInfoByUserName(t *testing.T) {
	err := UpdateStudentBasicInfoByUserName("jack", map[string]interface{}{
		"gender": 1,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestGetStudentBasicInfoByUserName(t *testing.T) {
	basicinfo, err := GetStudentBasicInfoByUserName("jack")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(basicinfo)
}

func TestAddStudyInfo(t *testing.T) {
	err := AddStudyInfo("jack")
	if err != nil {
		t.Error(err)
	}
}

func TestGetStudyInfoByUserName(t *testing.T) {
	studyinfo, err := GetStudyInfoByUserName("jack")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(studyinfo)
}

func TestUpdateStudyInfoByUserName(t *testing.T) {
	err := UpdateStudyInfoByUserName("jack", map[string]interface{}{
		"school_name": "ecnu",
	})
	if err != nil {
		t.Error(err)
	}
}
