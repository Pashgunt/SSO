package models

import (
	"gorm.io/gorm"
	"time"
)

type App struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Secret    string
	IsActual  int8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
