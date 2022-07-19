package server

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "app/auth/pkg/proto/gen"
)

func (s *Server) List(ctx context.Context,
	req *pb.ListRequest) (*pb.ListResponse, error) {
	answer, err := s.s.List(ctx, req.GetLastId(), req.GetLimit())
	if err != nil {
		return &pb.ListResponse{},
			s.getError(err)
	}

	var users []*pb.User
	for _, user := range answer {
		users = append(users, user.Proto())
	}

	return &pb.ListResponse{
		Users: users,
	}, nil
}

func (s *Server) User(ctx context.Context,
	req *pb.UserRequest) (*pb.UserResponse, error) {
	answer, err := s.s.User(ctx, req.GetId())
	if err != nil {
		return &pb.UserResponse{},
			s.getError(err)
	}

	return &pb.UserResponse{
		User: answer.Proto(),
	}, nil
}

func (s *Server) Create(ctx context.Context,
	req *pb.CreateRequest) (*emptypb.Empty, error) {
	id, err := s.s.Create(ctx,
		req.GetLogin(), req.GetPassword())
	if err != nil {
		return &emptypb.Empty{}, s.getError(err)
	}

	s.event <- &pb.UpdateStreamResponse{
		Action: pb.UpdateStreamResponse_ADD,
		User: &pb.User{
			Id:    id,
			Name:  req.GetLogin(),
			Login: req.GetLogin(),
			Email: "",
		},
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) Delete(ctx context.Context,
	req *pb.DeleteRequest) (*emptypb.Empty, error) {
	err := s.s.Delete(ctx, req.GetId())
	if err != nil {
		return &emptypb.Empty{}, s.getError(err)
	}

	s.event <- &pb.UpdateStreamResponse{
		Action: pb.UpdateStreamResponse_DELETE,
		User: &pb.User{
			Id: req.GetId(),
		},
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) ResetPassword(ctx context.Context,
	req *pb.ResetPasswordRequest) (*emptypb.Empty, error) {
	log.Println(req.GetNewPassword())
	log.Println(req.GetUserId())
	err := s.s.ChangeUserPassword(ctx,
		req.GetNewPassword(), req.GetNewPassword(), true, req.GetUserId())
	if err != nil {
		return &emptypb.Empty{}, s.getError(err)
	}

	return &emptypb.Empty{}, nil
}
