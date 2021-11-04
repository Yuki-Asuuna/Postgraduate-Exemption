package database

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/utils/mysql"
	"github.com/sirupsen/logrus"
)

func AddImage(ID int64, path string) error {
	image := &Image{
		ImageID:   ID,
		ImagePath: path,
	}
	if err := mysql.GetMySQLClient().Create(&image).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddImage Failed, err= %v", err)
		return err
	}
	return nil
}

func GetImages(IDs []int64) (map[int64]string, error) {
	ret := make(map[int64]string, 0)
	Images := make([]*Image, 0)
	err := mysql.GetMySQLClient().Where("image_id in (?)", IDs).Find(&Images).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"GetImages Failed, err= %v", err)
		return nil, err
	}
	for _, image := range Images {
		ret[image.ImageID] = image.ImagePath
	}
	return ret, nil
}
