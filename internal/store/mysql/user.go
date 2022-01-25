package mysql

import (
	"gorm.io/gorm"
	"virgo/internal/model"
)

type users struct {
	db *gorm.DB
}

func newUsers(ds *datastore) *users {
	return &users{ds.db}
}

func (u *users) Create(user *model.User) error {
	return u.db.Create(user).Error
}

func (u *users) Update(user *model.User) error  {
	return u.db.Save(user).Error
}