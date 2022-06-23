package models

import pb "app/access/pkg/proto/gen"

type Service struct {
	ID      uint64
	Name    string
	Methods []*Method
}

func (s *Service) Proto() *pb.Service {
	var methods []*pb.Method
	for _, method := range s.Methods {
		methods = append(methods, method.Proto())
	}

	return &pb.Service{
		Id:      s.ID,
		Name:    s.Name,
		Methods: methods,
	}
}

type Method struct {
	ID   uint64
	Name string
}

func (m *Method) Proto() *pb.Method {
	return &pb.Method{
		Id:   m.ID,
		Name: m.Name,
	}
}
