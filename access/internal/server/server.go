package server

import (
	"app/access/internal/service"
	pb "app/access/pkg/proto/gen"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	s service.Service
	pb.UnimplementedAccessServer
}

func New(s service.Service) *Server {
	return &Server{
		s: s,
	}
}

// func (s *Server) UserAccesses(ctx context.Context, req *pb.UserAccessesRequest) (*pb.UserAccessesResponse, error) {

// 	if req.GetId() == 0 {
// 		return &pb.UserAccessesResponse{}, status.Errorf(codes.Aborted, "invalid input data")
// 	}

// 	answer, err := s.s.UserAccesses(ctx, req.Id)
// 	if err != nil {
// 		return &pb.UserAccessesResponse{}, err
// 	}

// 	var accesses []*pb.AccessUser
// 	for service, method := range answer {
// 		for _, method_name := range method {
// 			accesses = append(accesses, &pb.AccessUser{Service: service, Method: method_name})
// 		}
// 	}

// 	return &pb.UserAccessesResponse{Access: accesses}, nil
// }

func (s *Server) List(ctx context.Context, _ *emptypb.Empty) (*pb.ListResponse, error) {
	return &pb.ListResponse{}, nil
}

func (s *Server) Group(ctx context.Context, req *pb.GroupRequest) (*pb.GroupResponse, error) {
	return &pb.GroupResponse{}, nil
}

func (s *Server) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Server) UpdateGroup(ctx context.Context, req *pb.UpdateGroupRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Server) Users(ctx context.Context, req *pb.UsersRequest) (*pb.UsersResponse, error) {
	return &pb.UsersResponse{}, nil
}

func (s *Server) AddUser(ctx context.Context, req *pb.AddUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Server) RemoveUser(ctx context.Context, req *pb.RemoveUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Server) ListServices(ctx context.Context, _ *emptypb.Empty) (*pb.ListServicesResponse, error) {
	return &pb.ListServicesResponse{}, nil
}

func (s *Server) AddMethod(ctx context.Context, req *pb.AddMethodRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Server) RemoveMethod(ctx context.Context, req *pb.RemoveMethodRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
