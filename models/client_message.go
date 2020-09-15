package models

import (
	"gorm.io/gorm"
	"time"
)

type ClientMessage struct {
	ID        uint `gorm:"primary_key"`
	ClientId  string
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
