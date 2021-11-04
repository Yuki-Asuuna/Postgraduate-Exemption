package database

import (
	"Postgraduate-Exemption/utils/mysql"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	mysql.MysqlInit()
	m.Run()
	os.Exit(0)
}

func TestGetImage(t *testing.T) {
	//t.Skip()
	ret, err := GetImages([]int64{244115596130648064, 244116453656100864, 0})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ret)
}
