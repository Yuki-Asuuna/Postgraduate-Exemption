package database

type StudyInfo struct {
	UserName        string
	StudyInfoID     int64
	SchoolName      string
	Writing         string
	AwardPunishment string
	Cheating        string
}

func (StudyInfo) TableName() string {
	return "study_info"
}
