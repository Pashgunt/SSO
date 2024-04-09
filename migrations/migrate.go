package migrations

import (
	"gorm.io/gorm"
	"os/user"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&user.User{})
	if err != nil {
		return err
	}

	return nil
}
