package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"app/auth/internal/service"
	pb "app/auth/pkg/proto/gen"
)

type Server struct {
	s     service.Service
	u     *Updater
	event chan *pb.UpdateStreamResponse

	pb.UnimplementedAuthServer
}

func New(s service.Service) *Server {

	events := make(chan *pb.UpdateStreamResponse)

	updater := NewUpdater(events)

	serv := &Server{
		s:     s,
		u:     updater,
		event: events,
	}

	return serv
}

func (s *Server) getError(err error) error {
	return status.Error(codes.Aborted, err.Error())
}

func (s *Server) SignIn(ctx context.Context,
	req *pb.SignInRequest) (*pb.SignInResponse, error) {
	token, user, access, err := s.s.
		SignIn(ctx, req.GetLogin(), req.GetPassword())
	if err != nil {
		return &pb.SignInResponse{}, s.getError(err)
	}

	userAccess := make(map[string]*pb.Methods)
	for service, methods := range access {
		userAccess[service] = &pb.Methods{Name: methods}
	}

	return &pb.SignInResponse{
		Token:  token,
		User:   user.Proto(),
		Access: userAccess,
	}, nil
}
