package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid();"`
	Email     string    `gorm:"uniqueIndex:user_email_unique_index,sort:desc"`
	PassHash  string
	IsActual  int8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
