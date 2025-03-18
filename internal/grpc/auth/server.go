package auth

import (
	"context"
	pant_sso_v1 "github.com/Pashgunt/Sso-Protobuf-Golang/gen/go/sso"
	"google.golang.org/grpc"
)

type Auth interface {
	Login( // метод для аутентификации пользователя. Он принимает email, пароль и UUID приложения,
		// возвращает токен и UUID пользователя, а также возможную ошибку.
		ctx context.Context,
		email string,
		password string,
		appUuid string,
	) (token string, userUuid string, err error)
	RegisterNewUser( //метод для регистрации нового пользователя. Он принимает email и пароль,
		// возвращает UUID пользователя и возможную ошибку.
		ctx context.Context,
		email string,
		password string,
	) (userUuid string, err error)
	ForgotPassword(
		ctx context.Context,
		email string,
		appUuid string,
	) (token string, userUuid string, err error)
}

// Она встраивает pant_sso_v1.UnimplementedAuthServer,
// который является сгенерированным базовым классом для сервера в gRPC.
// Это необходимо для обеспечения совместимости с интерфейсами, определёнными в .proto файле.
type ServerAuthApi struct {
	pant_sso_v1.UnimplementedAuthServer
	auth Auth // реализация методов аутентификации.
}

// Создаёт экземпляр ServerAuthApi, в который передаётся реализация интерфейса Auth (параметр auth).
// Регистрирует сервер аутентификации в gRPC-сервере.
// pant_sso_v1.RegisterAuthServer — это автоматически сгенерированная gRPC-функция, которая настраивает сервер для обработки запросов,
// определённых в gRPC-сервисе (в .proto файле).
// Она ожидает, что передаваемый сервер будет реализовывать методы, такие как Login и RegisterNewUser,
// которые описаны в интерфейсе Auth.
func RegisterServerApiHandler(server *grpc.Server, auth Auth) {
	pant_sso_v1.RegisterAuthServer(server, &ServerAuthApi{auth: auth})
}

//RegisterServerApiHandler помогает зарегистрировать сервер аутентификации в gRPC,
//привязав реализацию методов (Login, RegisterNewUser) к gRPC-серверу.
//Это позволяет серверу обрабатывать запросы, приходящие от клиентов,
//и вызывать соответствующие методы в реализации аутентификации.
