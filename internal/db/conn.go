package db

import (
	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	db   *gorm.DB
	err  error
	once sync.Once
)

func DB() *gorm.DB {
	once.Do(func() {
		driver := viper.GetString("db.driver")
		dsn := viper.GetString("db.dsn")
		db, err = connect(driver, dsn)
		if err != nil {
			panic(err)
		}
	})
	return db
}

func connect(driver, dsn string) (*gorm.DB, error) {
	switch driver {
	case "sqlite3", "sqlite":
		return connectSQLite(dsn)
	case "postgres":
		return connectPostgres(dsn)
	default:
		// use sqlite as default
		return connectSQLite(dsn)
	}
}

func connectSQLite(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		dsn = ":memory:"
	}
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}

func connectPostgres(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
