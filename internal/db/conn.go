package db

import (
	"github.com/glebarez/sqlite"
	"github.com/pplmx/blog-go/internal/model"
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

func NewDBConn() *gorm.DB {
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

func connect(driver, dsn string) (db *gorm.DB, err error) {
	switch driver {
	case "sqlite3", "sqlite":
		db, err = connectSQLite(dsn)
	case "postgres":
		db, err = connectPostgres(dsn)
	default:
		// use sqlite as default
		db, err = connectSQLite(dsn)
	}
	if err != nil {
		return nil, err
	}

	// migrate schema
	tables := []interface{}{
		&model.Post{},
	}
	err = db.AutoMigrate(tables...)
	if err != nil {
		return db, err
	}
	return db, nil
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
