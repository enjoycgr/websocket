package models

import (
	"gorm.io/gorm"
	"time"
)

type Group struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(50)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
