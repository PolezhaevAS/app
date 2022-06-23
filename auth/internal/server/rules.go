package server

import (
	"app/internal/service"

	"google.golang.org/grpc"
)

func Rules(s *grpc.ServiceDesc) *service.Service {
	service := service.New(s)

	// For example
	// service.AddOpenApi("List")
	service.AddOpenApi("SignIn")
	return service
}
