package authservice

import (
	"context"
	"errors"
	"fmt"
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

	//Выполняем запрос к БД: ищем пользователя по email.
	if err := a.checkCorrectOrmOperation(a.Gorm.Where("email = ?", email).First(&user)); err != nil {
		return "", "", errors.New(fmt.Sprintf("This email is not registered, %v", err.Error()))
	}

	//проверяет, совпадает ли хэш пароля с сохранённым значением.
	if !hasher.CheckPasswordHash(password, user.PassHash) {
		return "", "", errors.New("you entered an incorrect password")
	}

	// Преобразуем appUuid из строки в uuid.UUID.
	parseUuid, _ := uuid.Parse(appUuid)

	// Проверяем, существует ли такое приложение в БД.
	if err := a.checkCorrectOrmOperation(a.Gorm.Where("id = ?", parseUuid).First(&app)); err != nil {
		return "", "", errors.New(fmt.Sprintf("This app is not exists, %v", err.Error()))
	}

	// создаёт новый токен с учётом времени жизни TokenTtl.
	jwtToken, err := jwt.NewJwtToken(user, a.TokenTtl, app)

	if err != nil {
		return "", "", err
	}

	// Возвращаем: jwtToken — токен доступа.
	return jwtToken, user.ID.String(), nil
}

// регистрация нового пользователя)
func (a *Auth) RegisterNewUser(
	ctx context.Context,
	email string,
	password string,
) (userUuid string, err error) {
	hashedPassword, _ := hasher.MakeHashPassword(password) // хеширует пароль перед сохранением.
	user := models.User{                                   //Создаём объект пользователя с email, PassHash (захешированный пароль).
		Email:     email,
		PassHash:  hashedPassword,
		IsActual:  1,
		CreatedAt: time.Time{}, // Поля CreatedAt, UpdatedAt, DeletedAt задаются по умолчанию.
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
	errorCreate := a.Gorm.Create(&user).Error // вставляет пользователя в БД.

	return user.ID.String(), errorCreate
}

func (a *Auth) checkCorrectOrmOperation(result *gorm.DB) error { // Этот метод обрабатывает ошибки, возникающие при работе с БД:
	if result.Error == nil {
		return nil
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) { //Если ошибка — запись не найдена → возвращаем её.
		return result.Error
	}

	panic(result.Error) // Во всех остальных случаях вызываем panic, потому что это критическая ошибка (например, сбой подключения к БД).
}

func (a *Auth) ForgotPassword(
	ctx context.Context,
	email string,
	appUuid string,
) (token string, userUuid string, err error) {
	var user models.User
	var app models.App

	parseUuid, _ := uuid.Parse(appUuid)
	// Проверяем, существует ли такое приложение в БД.
	if err := a.checkCorrectOrmOperation(a.Gorm.Where("id = ?", parseUuid).First(&app)); err != nil {
		return "", "", errors.New(fmt.Sprintf("This app is not exists, %v", err.Error()))
	}

	// Проверяем, существует ли пользователь с таким email
	if err := a.checkCorrectOrmOperation(a.Gorm.Where("email = ?", email).First(&user)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "", fmt.Errorf("user with email %s not found", email)
		}
		return "", "", err // Если другая ошибка → возвращаем её
	}

	//TODO отправить письмо с ссылкой восстановления на почту

	return token, userUuid, fmt.Errorf("Проверьте письмо отправленное по адресу %s n", email)
}
