package mysql

import (
	"log"
	"testing"
	"time"
	"virgo/internal/model"
	"virgo/pkg/options"
)

func TestGetMySQLFactory(t *testing.T) {

	storeFactory, err := GetMySQLFactory(&options.MySQLOptions{
		Host: "localhost:3306",
		Username: "virgo",
		Password: "123456",
		Database: "virgo",
		LogLevel: 3,
	})
	if err != nil {
		log.Fatal(err)
	}

	userStore := storeFactory.Users()
	_ = userStore.Update(&model.User{
		Id:       1,
		Name:     "xuliang",
		Age:      27,
		Gender:   1,
		Email:    "grasswort@qq.com",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	})
}
