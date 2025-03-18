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

//Открывается соединение с PostgreSQL через gorm.Open().
//Используется gorm.io/driver/postgres.
//В качестве подключения (Conn) передаётся hs.handleServiceStructure.Psql.Db().
//Если соединение не удалось (err != nil), приложение панически завершается (panic(err)).
//Создаётся и возвращается структура authservice.Auth, которая содержит:
//Gorm — объект gorm.DB для работы с базой данных.
//Redis — клиент Redis (hs.handleServiceStructure.RedisApp.Client()).
//TokenTtl — время жизни токена.
//Log — логгер.

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
