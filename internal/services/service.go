package services

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sso/internal/config"
	authservice "sso/internal/services/auth"
)

type HandlerServices struct {
	handleServiceStructure *config.HandleServiceStructure
}

func NewHandlerServices(handleServiceStructure *config.HandleServiceStructure) *HandlerServices {
	return &HandlerServices{handleServiceStructure: handleServiceStructure}
}

func (hs *HandlerServices) MakeAuthService() *authservice.Auth {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: hs.handleServiceStructure.Psql.Db(),
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &authservice.Auth{
		Gorm:     gormDB,
		Redis:    hs.handleServiceStructure.RedisApp.Client(),
		TokenTtl: hs.handleServiceStructure.TokenTtl,
		Log:      hs.handleServiceStructure.Log,
	}
}
