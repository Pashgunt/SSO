package authservice

import (
	"context"
)

type Auth struct {
}

func (a *Auth) Login(
	ctx context.Context,
	email string,
	password string,
	appUuid string,
) (token string, err error) {
	return email, nil
}

func (a *Auth) RegisterNewUser(
	ctx context.Context,
	email string,
	password string,
) (userUuid string, err error) {
	return email, nil
}
