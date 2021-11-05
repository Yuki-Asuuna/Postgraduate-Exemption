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
	BirthLocationProvince  int64
	BirthLocationCity      int64
	BirthLocationCounty    int64
	NativeLocationProvince int64
	NativeLocationCity     int64
	NativeLocationCounty   int64
	ImageID                int64
}

func (StudentBasicInfo) TableName() string {
	return "student_basic_info"
}
