package api

type GetAccountInfoResponse struct {
	UserName    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
	Identity    int64  `json:"identity"`
	University  string `json:"university"`
	Major       string `json:"major"`
}

type GetStudentBasicInfoResponse struct {
	UserName               string `json:"username"`
	IdentityNumber         string `json:"identityNumber"`
	Name                   string `json:"name"`
	NamePinyin             string `json:"namePinYin"`
	MilitaryType           int64  `json:"militaryType"`
	PoliticalStatus        int64  `json:"politicalStatus"`
	Gender                 int64  `json:"gender"`
	MartialStatus          int64  `json:"martialStatus"`
	BirthLocationProvince  string `json:"birthLocationProvince"`
	BirthLocationCity      string `json:"birthLocationCity"`
	BirthLocationCounty    string `json:"birthLocationCounty"`
	NativeLocationProvince string `json:"nativeLocationProvince"`
	NativeLocationCity     string `json:"nativeLocationCity"`
	NativeLocationCounty   string `json:"nativeLocationCounty"`
	ImageID                int64  `json:"imageID"`
}

type GetProfileInfoResponse struct {
	UserName                  string `json:"username"`
	ProfileID                 int64 `json:"profileID"`
	ProfileLocationProvince   string `json:"profileLocationProvince"`
	ProfileLocationCity       string `json:"profileLocationCity"`
	ProfileLocationCounty     string `json:"profileLocationCounty"`
	ProfileAddress            string `json:"profileAddress"`
	PostCode                  string `json:"postCode"`
	ResidenceLocationProvince string `json:"residenceLocationProvince"`
	ResidenceLocationCity     string `json:"residenceLocationCity"`
	ResidenceLocationCounty   string `json:"residenceLocationCounty"`
	ResidenceAddress          string `json:"residenceAddress"`
}
