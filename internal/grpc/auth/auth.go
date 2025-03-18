package auth

import (
	"context"
	pant_sso_v1 "github.com/Pashgunt/Sso-Protobuf-Golang/gen/go/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerAuthApi) Login(
	ctx context.Context,
	request *pant_sso_v1.LoginRequest,
) (*pant_sso_v1.LoginResponse, error) {

	if err := s.validateLogin(request); err != nil { //проверку данных запроса
		return nil, err
	}

	token, userUuid, err := s.auth.Login( // Если валидация прошла успешно, вызывается метод Login
		ctx,
		request.GetEmail(),
		request.GetPassword(),
		request.GetAppUuid(),
	)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pant_sso_v1.LoginResponse{
		Uuid:     userUuid,
		AppUuid:  request.GetAppUuid(),
		JwtToken: token,
	}, nil
}

func (s *ServerAuthApi) Register(
	ctx context.Context,
	request *pant_sso_v1.RegisterRequest,
) (*pant_sso_v1.RegisterResponse, error) {

	if err := s.validateRegister(request); err != nil { //проверяет корректность данных регистрации
		return nil, err
	}

	//Если данные валидны, вызывается метод RegisterNewUser из объекта auth (реализация интерфейса Auth),
	userId, err := s.auth.RegisterNewUser(ctx, request.GetEmail(), request.GetPassword())

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pant_sso_v1.RegisterResponse{Uuid: userId}, nil
}

func (s *ServerAuthApi) ForgotPassword(
	ctx context.Context,
	request *pant_sso_v1.ForgotPasswordRequest,
) (*pant_sso_v1.ForgotPasswordResponse, error) {

	if err := s.validateForgotPassword(request); err != nil { //проверку данных запроса
		return nil, err
	}

	token, userUuid, err := s.auth.ForgotPassword(
		ctx,
		request.GetEmail(),
		request.GetAppUuid(),
	)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pant_sso_v1.ForgotPasswordResponse{
		Uuid:     userUuid,
		AppUuid:  request.GetAppUuid(),
		JwtToken: token,
	}, nil
}
