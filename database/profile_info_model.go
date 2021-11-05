package database

type ProfileInfo struct {
	UserName                  string
	ProfileID                 int64
	ProfileLocationProvince   int64
	ProfileLocationCity       int64
	ProfileLocationCounty     int64
	ProfileAddress            string
	PostCode                  string
	ResidenceLocationProvince int64
	ResidenceLocationCity     int64
	ResidenceLocationCounty   int64
}

func (ProfileInfo) TableName() string {
	return "profile_info"
}
