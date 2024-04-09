package auth

import (
	"context"
	pant_sso_v1 "github.com/Pashgunt/Sso-Protobuf-Golang/gen/go/sso"
	"google.golang.org/grpc"
)

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appUuid string,
	) (token string, err error)
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
	token, _ := s.auth.Login(
		ctx,
		request.GetEmail(),
		request.GetPassword(),
		request.GetAppUuid(),
	)

	return &pant_sso_v1.LoginResponse{
		UserUuid: "123123",
		AppUuid:  "1231231",
		JwtToken: token,
	}, nil
}

func (s *serverAuthApi) Register(
	ctx context.Context,
	in *pant_sso_v1.RegisterRequest,
) (*pant_sso_v1.RegisterResponse, error) {
	panic("implement me register")
}
