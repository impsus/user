package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConn() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	if os.Getenv("ENGINE") == "MYSQL" {
		db, err = gorm.Open(mysql.Open(os.Getenv("MYSQL_DB_CONN_STR")))
	}

	if os.Getenv("ENGINE") == "SQLITE" {
		db, err = gorm.Open(sqlite.Open("test.db"))
	}

	if err != nil {
		return nil, err
	}
	return db, nil
}
