package database

type ProfileInfo struct {
	UserName                  string
	ProfileID                 int64
	ProfileLocationProvince   string
	ProfileLocationCity       string
	ProfileLocationCounty     string
	ProfileAddress            string
	PostCode                  string
	ResidenceLocationProvince string
	ResidenceLocationCity     string
	ResidenceLocationCounty   string
	ResidenceAddress          string
}

func (ProfileInfo) TableName() string {
	return "profile_info"
}
