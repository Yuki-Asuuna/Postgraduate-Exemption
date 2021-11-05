package database

type FamilyMember struct {
	ProfileID    int64
	Name         string
	Relationship string
	PhoneNumber  string
	Address      string
}

func (FamilyMember) TableName() string {
	return "family_member"
}
