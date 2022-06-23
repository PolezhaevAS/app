package server

import (
	"app/access/internal/models"
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

func (s *Server) List(ctx context.Context, _ *emptypb.Empty) (*pb.ListResponse, error) {
	answer, err := s.s.List(ctx)
	if err != nil {
		return nil, err
	}

	resp := &pb.ListResponse{}
	groups := []*pb.Group{}
	for _, group := range answer {
		groups = append(groups, group.Proto())
	}

	resp.Groups = groups

	return resp, nil
}

func (s *Server) Group(ctx context.Context, req *pb.GroupRequest) (*pb.GroupResponse, error) {
	answer, err := s.s.Group(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GroupResponse{Group: answer.Proto()}, nil
}

func (s *Server) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*emptypb.Empty, error) {
	_, err := s.s.CreateGroup(ctx, req.GetName(), req.GetDesc())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) UpdateGroup(ctx context.Context, req *pb.UpdateGroupRequest) (*emptypb.Empty, error) {
	err := s.s.UpdateGroup(ctx, &models.Group{
		ID:   req.GetId(),
		Name: req.GetName(),
		Desc: req.GetDesc(),
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*emptypb.Empty, error) {
	err := s.s.DeleteGroup(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) Users(ctx context.Context, req *pb.UsersRequest) (*pb.UsersResponse, error) {
	answer, err := s.s.Users(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.UsersResponse{Users: answer}, nil
}

func (s *Server) AddUser(ctx context.Context, req *pb.AddUserRequest) (*emptypb.Empty, error) {
	err := s.s.AddUser(ctx, req.GetGroupId(), req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) RemoveUser(ctx context.Context, req *pb.RemoveUserRequest) (*emptypb.Empty, error) {
	err := s.s.RemoveUser(ctx, req.GetGroupId(), req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) ListServices(ctx context.Context, _ *emptypb.Empty) (*pb.ListServicesResponse, error) {
	answer, err := s.s.ListService(ctx)
	if err != nil {
		return nil, err
	}
	var services []*pb.Service
	for _, s := range answer {
		services = append(services, s.Proto())
	}

	return &pb.ListServicesResponse{Services: services}, nil
}

func (s *Server) AddMethod(ctx context.Context, req *pb.AddMethodRequest) (*emptypb.Empty, error) {
	err := s.s.AddMethod(ctx, req.GetGroupId(), req.GetMethodId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) RemoveMethod(ctx context.Context, req *pb.RemoveMethodRequest) (*emptypb.Empty, error) {
	err := s.s.RemoveMethod(ctx, req.GetGroupId(), req.GetMethodId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
