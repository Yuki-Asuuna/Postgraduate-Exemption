package database

type ContactInfo struct {
	UserName        string
	PhoneNumber     string
	FixedLineNumber string
	Address         string
	PostCode        string
}

func (ContactInfo) TableName() string {
	return "contact_info"
}
