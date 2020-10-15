package models

import (
	"gorm.io/gorm"
	"time"
)

type GroupClient struct {
	ID        uint   `gorm:"primary_key"`
	ClientID  string `gorm:"type:varchar(50)"`
	GroupID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
