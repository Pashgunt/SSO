package auth

import (
	"context"
	pant_sso_v1 "github.com/Pashgunt/Sso-Protobuf-Golang/gen/go/sso"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appUuid string,
	) (token string, userUuid string, err error)
	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (userUuid string, err error)
}

type serverAuthApi struct {
	pant_sso_v1.UnimplementedAuthServer
	auth Auth
}

func RegisterServerApiHandler(server *grpc.Server, auth Auth) {
	pant_sso_v1.RegisterAuthServer(server, &serverAuthApi{auth: auth})
}

func (s *serverAuthApi) Login(
	ctx context.Context,
	request *pant_sso_v1.LoginRequest,
) (*pant_sso_v1.LoginResponse, error) {
	if err := s.validateLogin(request); err != nil {
		return nil, err
	}

	token, userUuid, err := s.auth.Login(
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

func (s *serverAuthApi) Register(
	ctx context.Context,
	request *pant_sso_v1.RegisterRequest,
) (*pant_sso_v1.RegisterResponse, error) {
	if err := s.validateRegister(request); err != nil {
		return nil, err
	}

	userId, err := s.auth.RegisterNewUser(ctx, request.GetEmail(), request.GetPassword())

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pant_sso_v1.RegisterResponse{Uuid: userId}, nil
}

func (s *serverAuthApi) validateRegister(
	request *pant_sso_v1.RegisterRequest,
) error {
	validate := validator.New()

	err := validate.Var(request.GetEmail(), "required,email")

	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	err = validate.Var(request.GetPassword(), "required,min=8")

	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	return nil
}

func (s *serverAuthApi) validateLogin(
	request *pant_sso_v1.LoginRequest,
) error {
	validate := validator.New()

	err := validate.Var(request.GetEmail(), "required,email")

	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	err = validate.Var(request.GetPassword(), "required,min=8")

	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	err = validate.Var(request.GetAppUuid(), "required")

	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	return nil
}
