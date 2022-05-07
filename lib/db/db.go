package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConn() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_DB_CONN_STR")))
	if err != nil {
		return nil, err
	}
	return db, nil
}
