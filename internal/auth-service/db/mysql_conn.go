package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConn() (*gorm.DB, error) {
	dsn := string(os.Getenv("MYSQL_URI"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
