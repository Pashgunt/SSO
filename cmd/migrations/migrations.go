package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sso/cmd/global"
	psqlapp "sso/internal/app/database/psql"
	"sso/internal/domain/models"
)

func main() {
	cfg, logger := global.InitGlobalStructures()
	logger.Info("Start Migrations")
	psqlConnection := psqlapp.NewPsqlApp(
		logger,
		cfg.PSQL.Username,
		cfg.PSQL.Dbname,
		cfg.PSQL.Password,
		cfg.PSQL.Host,
		cfg.PSQL.Port,
	)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: psqlConnection.Db(),
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err := gormDB.AutoMigrate(&models.User{}, &models.App{}); err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			logger.Error(fmt.Sprintf("Error processing Migrations %s", err))
		}

		logger.Info("Success end Migrations")

		psqlConnection.Stop()
	}()
}
