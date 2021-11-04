package database

type Image struct {
	ImageID   int64
	ImagePath string
}

func (Image) TableName() string {
	return "image"
}
