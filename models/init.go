package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Migrate() {
	db.AutoMigrate(&User{})
}

func InitDB(user, password, addr, dbName string) (err error) {
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True",
		user,
		password,
		addr,
		dbName)), &gorm.Config{})

	if err != nil {
		return err
	}
	return nil
}
