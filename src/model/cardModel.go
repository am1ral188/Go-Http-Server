package model

import (
	"awesomeProject/src/tools"
	"gorm.io/gorm"
)

type CardModel struct {
	DB *gorm.DB
}
type card struct {
	ID     int64  `gorm:"primaryKey;autoIncrement"`
	UserID string `gorm:"not null"`
	Hash   string `gorm:"not null"`
}

func (m CardModel) Connect(host, username, password, dbName string) (*gorm.DB, error) {
	return tools.Connect(host, username, password, dbName)
}
func (r CardModel) NewCard() card {
	return card{}
}
