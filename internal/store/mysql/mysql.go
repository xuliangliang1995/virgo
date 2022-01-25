package mysql

import (
	"gorm.io/gorm"
	"virgo/internal/store"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Users() store.UserStore  {
	return newUsers(ds)
}

func (ds *datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
