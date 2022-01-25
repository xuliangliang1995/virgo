package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
	"time"
	"virgo/internal/store"
	"virgo/pkg/options"
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


var (
	mysqlFactory store.Factory
	once sync.Once
)

func GetMySQLFactory(opts *options.MySQLOptions) (store.Factory, error)  {
	if opts == nil && mysqlFactory == nil {
		return nil, fmt.Errorf("failed to get mysql store factory")
	}

	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
			opts.Username,
			opts.Password,
			opts.Host,
			opts.Database,
			true,
			"Local")

		dbIns, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.New(log.Default(), logger.Config{
				LogLevel: logger.LogLevel(opts.LogLevel),
				Colorful:                  true,
				IgnoreRecordNotFoundError: false,
				SlowThreshold:             10 * time.Second,
			}),
		})

		mysqlFactory = &datastore{dbIns}
	})


	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, err

}
