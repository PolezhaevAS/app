package server

import (
	pb "app/auth/pkg/proto/gen"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ChangeUser(ctx context.Context,
	req *pb.ChangeUserRequest) (*emptypb.Empty, error) {
	err := s.s.ChangeUser(ctx,
		req.GetName(), req.GetLogin(), req.GetEmail())
	if err != nil {
		return &emptypb.Empty{}, s.getError(err)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) ChangeUserPassword(ctx context.Context,
	req *pb.ChangeUserPasswordRequest) (*emptypb.Empty, error) {
	err := s.s.ChangeUserPassword(ctx,
		req.GetOldPassword(), req.GetNewPassword(), false)
	if err != nil {
		return &emptypb.Empty{}, s.getError(err)
	}

	return &emptypb.Empty{}, nil
}
