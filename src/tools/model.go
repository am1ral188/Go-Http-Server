package tools

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(host, username, password, dbName string) (*gorm.DB, error) {
	dsn := username + ":" + password + "@tcp(" + host + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return DB, err
	}
	return DB, nil
}
