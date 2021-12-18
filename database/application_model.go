package database

import "time"

type Application struct {
	ApplicationID int64
	UserName      string
	University    string
	Major         string
	IsAdmitted    int64
	IsConfirmed   int64
	SubmitTime    time.Time
}

func (Application) TableName() string {
	return "application"
}
