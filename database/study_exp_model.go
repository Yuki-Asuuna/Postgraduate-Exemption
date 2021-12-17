package database

type StudyExp struct {
	ExperienceID int64
	StudyInfoID  int64
	Interval     string
	WorkPlace    string
}

func (StudyExp) TableName() string {
	return "study_exp"
}
