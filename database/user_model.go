package database

import "time"

type User struct {
	UserName    string
	Password    string
	Identity    int64
	PhoneNumber string
	CreateTime  time.Time
	UpdateTime  time.Time
}

func (User) TableName() string {
	return "user"
}
