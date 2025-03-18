package auth

import (
	"fmt"
	pant_sso_v1 "github.com/Pashgunt/Sso-Protobuf-Golang/gen/go/sso"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	validate = validator.New()
)

func (s *ServerAuthApi) validateRegister(
	request *pant_sso_v1.RegisterRequest,
) error {

	err := validate.Var(request.GetEmail(), "required,email")

	if err != nil {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("Email validation error: enter the correct email address. \n %v", err.Error()))
	}

	err = validate.Var(request.GetPassword(), "required,min=8")

	if err != nil {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("Password validation error: The minimum password length is 8 characters. \n %v", err.Error()))
	}

	return nil
}

func (s *ServerAuthApi) validateLogin(
	request *pant_sso_v1.LoginRequest,
) error {

	err := validate.Var(request.GetEmail(), "required,email")

	if err != nil {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("Email validation error: enter the correct email address. \n %v", err.Error()))
	}

	err = validate.Var(request.GetPassword(), "required,min=8")

	if err != nil {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("Password validation error: The minimum password length is 8 characters. \n %v", err.Error()))
	}

	err = validate.Var(request.GetAppUuid(), "required")

	if err != nil {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("Enter the correct AppId. \n %v", err.Error()))
	}

	return nil
}

func (s *ServerAuthApi) validateForgotPassword(request *pant_sso_v1.ForgotPasswordRequest) error {

	err := validate.Var(request.GetEmail(), "required,email")

	if err != nil {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("Email validation error: enter the correct email address. \n %v", err.Error()))
	}

	err = validate.Var(request.GetAppUuid(), "required")

	if err != nil {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("Enter the correct AppId. \n %v", err.Error()))
	}

	return nil
}
