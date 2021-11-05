package database

type Agreement struct {
	UserName        string
	HasAgreedNotice int64
	HasAgreedHonest int64
}

func (Agreement) TableName() string {
	return "agreement"
}
