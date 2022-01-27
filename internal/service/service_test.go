package service

import (
	"log"
	"testing"
	"time"
	"virgo/internal/model"
	"virgo/internal/store/mysql"
	"virgo/pkg/options"
)

func TestUserService(t *testing.T) {
	storeFactory, _ := mysql.GetMySQLFactory(&options.MySQLOptions{
		Host: "localhost:3306",
		Username: "virgo",
		Password: "123456",
		Database: "virgo",
		LogLevel: 3,
	})

	srv := NewService(storeFactory)

	err := srv.Users().Update(&model.User{
		Id:       1,
		Name:     "xuliangliang",
		Age:      27,
		Gender:   1,
		Email:    "grasswort@qq.com",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	})
	if err != nil {
		log.Fatal(err)
	}
}
