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
	validate := validator.New()

	err := validate.Var(request.GetEmail(), "required,email")

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = validate.Var(request.GetPassword(), "required,min=8")

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = validate.Var(request.GetAppUuid(), "required")

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token, userUuid, _ := s.auth.Login(
		ctx,
		request.GetEmail(),
		request.GetPassword(),
		request.GetAppUuid(),
	)

	return &pant_sso_v1.LoginResponse{
		UserUuid: userUuid,
		AppUuid:  request.GetAppUuid(),
		JwtToken: token,
	}, nil
}

func (s *serverAuthApi) Register(
	ctx context.Context,
	request *pant_sso_v1.RegisterRequest,
) (*pant_sso_v1.RegisterResponse, error) {
	validate := validator.New()

	err := validate.Var(request.GetEmail(), "required,email")

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = validate.Var(request.GetPassword(), "required,min=8")

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userId, _ := s.auth.RegisterNewUser(ctx, request.GetEmail(), request.GetPassword())

	return &pant_sso_v1.RegisterResponse{UserUuid: userId}, nil
}
