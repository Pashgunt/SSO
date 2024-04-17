package authservice

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log/slog"
	"sso/internal/domain/models"
	"sso/internal/lib/hasher"
	"sso/internal/lib/jwt"
	"time"
)

type Auth struct {
	Gorm     *gorm.DB
	Redis    *redis.Client
	TokenTtl time.Duration
	Log      *slog.Logger
}

func (a *Auth) Login(
	ctx context.Context,
	email string,
	password string,
	appUuid string,
) (token string, userUuid string, err error) {
	var user models.User
	var app models.App

	if err := a.checkCorrectOrmOperation(a.Gorm.Where("email = ?", email).First(&user)); err != nil {
		return "", "", err
	}

	if !hasher.CheckPasswordHash(password, user.PassHash) {
		return "", "", errors.New("IncorrectPassword")
	}

	parseUuid, _ := uuid.Parse(appUuid)

	if err := a.checkCorrectOrmOperation(a.Gorm.Where("id = ?", parseUuid).First(&app)); err != nil {
		return "", "", err
	}

	jwtToken, err := jwt.NewJwtToken(user, a.TokenTtl, app)

	if err != nil {
		return "", "", err
	}

	return jwtToken, user.ID.String(), nil
}

func (a *Auth) RegisterNewUser(
	ctx context.Context,
	email string,
	password string,
) (userUuid string, err error) {
	hashedPassword, _ := hasher.MakeHashPassword(password)
	user := models.User{
		Email:     email,
		PassHash:  hashedPassword,
		IsActual:  1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
	errorCreate := a.Gorm.Create(&user).Error

	return user.ID.String(), errorCreate
}

func (a *Auth) checkCorrectOrmOperation(result *gorm.DB) error {
	if result.Error == nil {
		return nil
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	panic(result.Error)
}
