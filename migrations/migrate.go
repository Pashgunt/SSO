package migrations

import (
	"gorm.io/gorm"
	"sso/internal/domain/models"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.App{})

	if err != nil {
		return err
	}

	return nil
}
