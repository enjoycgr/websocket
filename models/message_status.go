package models

import (
	"gorm.io/gorm"
	"time"
)

type MessageStatus struct {
	ID        uint `gorm:"primary_key"`
	Read      uint8
	ClientID  uint
	MessageID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
