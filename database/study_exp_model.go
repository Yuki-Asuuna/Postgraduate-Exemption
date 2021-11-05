package database

type StudyExp struct {
	StudyInfoID int64
	Interval    string
	WorkPlace   string
}

func (StudyExp) TableName() string {
	return "study_exp"
}
