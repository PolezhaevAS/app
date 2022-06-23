package server

import (
	"app/auth/internal/models"
	"app/auth/internal/service"
	pb "app/auth/pkg/proto/gen"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	s service.Service
	pb.UnimplementedAuthServer
}

func New(s service.Service) *Server {
	return &Server{
		s: s,
	}
}

func (s *Server) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	token, user, err := s.s.Token(ctx, req.GetLogin(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	return &pb.SignInResponse{
		Token: token,
		User:  user.Proto(),
	}, nil
}

func (s *Server) User(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user, err := s.s.User(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{User: user.Proto()}, nil
}

func (s *Server) List(ctx context.Context, _ *emptypb.Empty) (*pb.ListResponse, error) {
	list, err := s.s.List(ctx)
	if err != nil {
		return nil, err
	}

	var users []*pb.User
	for _, user := range list {
		users = append(users, user.Proto())
	}

	return &pb.ListResponse{Users: users}, nil
}

func (s *Server) Create(ctx context.Context, req *pb.CreateRequest) (*emptypb.Empty, error) {
	err := s.s.Create(ctx, req.GetName(), req.GetLogin(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) Update(ctx context.Context, req *pb.UpdateRequest) (*emptypb.Empty, error) {
	err := s.s.Update(ctx, models.UserFromProto(req.User))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	err := s.s.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
