package models

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID        uint   `gorm:"primary_key"`
	Message   string `gorm:"type:varchar(255)"`
	Sender    string `gorm:"type:varchar(50)"`
	Owner     string `gorm:"type:tinyint(1)"`
	OwnerID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
