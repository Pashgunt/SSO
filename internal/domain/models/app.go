package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type App struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid();"`
	Name      string
	Secret    string
	IsActual  int8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
