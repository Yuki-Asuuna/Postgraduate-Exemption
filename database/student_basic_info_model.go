package database

type StudentBasicInfo struct {
	UserName               string
	IdentityNumber         string
	Name                   string
	NamePinyin             string
	MilitaryType           int64
	PoliticalStatus        int64
	Gender                 int64
	MartialStatus          int64
	BirthLocationProvince  string
	BirthLocationCity      string
	BirthLocationCounty    string
	NativeLocationProvince string
	NativeLocationCity     string
	NativeLocationCounty   string
	ImageID                int64
}

func (StudentBasicInfo) TableName() string {
	return "student_basic_info"
}
