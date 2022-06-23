package service

import (
	access "app/access/pkg/proto/gen"

	"google.golang.org/grpc"
)

type Service struct {
	service *grpc.ServiceDesc
	openApi []string
}

func New(service *grpc.ServiceDesc) *Service {
	return &Service{service: service}
}

func (s *Service) AddOpenApi(methodName string) {
	s.openApi = append(s.openApi, methodName)
}

func (s *Service) OpenApi(methodName string) bool {
	for _, m := range s.openApi {
		if m == methodName {
			return true
		}
	}
	return false
}

func (s *Service) Name() string {
	return s.service.ServiceName
}

func Services() (s []grpc.ServiceDesc) {

	s = append(s, access.Access_ServiceDesc)

	return
}
