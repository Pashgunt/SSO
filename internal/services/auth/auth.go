package authservice

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log/slog"
	"sso/internal/domain/models"
	"sso/internal/lib/hasher"
	"sso/internal/lib/jwt"
	"strconv"
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
	hashedPassword, _ := hasher.MakeHashPassword(password)
	var user models.User
	a.Gorm.Model(&models.User{
		Email:    email,
		PassHash: hashedPassword,
	}).Take(&user)
	jwtToken, err := jwt.NewJwtToken(user, a.TokenTtl, models.App{
		ID:        1,
		Name:      "Test App",
		Secret:    "TestApp",
		IsActual:  1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	})
	if err != nil {
		return "", "", err
	}
	return jwtToken, strconv.Itoa(int(user.ID)), nil
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

	a.Gorm.Create(&user)

	return strconv.Itoa(int(user.ID)), nil
}
