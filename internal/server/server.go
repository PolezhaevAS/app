package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"app/internal/server/auth"
)

// The Config represents gRPC server configurations.
type Config struct {
	Network string `json:"network" yaml:"network" toml:"network" mapstructure:"network"` //nolint
	Bind    string `json:"bind" yaml:"bind" toml:"bind" mapstructure:"bind"`             //nolint
}

func NewConfig() *Config {
	return &Config{}
}

type Server struct {
	// TCP socket for incoming server connection.
	l net.Listener

	// gRPC-server instance.
	grpc *grpc.Server

	// wait group for background workers, gracefully shutdown etc.
	wg sync.WaitGroup

	// global gRPC-server config.
	cfg *Config

	// jwt-auth instance.
	auth *auth.Auth
}

func New(conf *Config, auth *auth.Auth) (s *Server, err error) {
	s = new(Server)

	s.cfg = conf

	if s.l, err = net.Listen(conf.Network, conf.Bind); err != nil {
		return nil, err
	}

	// create a new gRPC server.
	s.grpc = grpc.NewServer(
		grpc.StreamInterceptor(auth.StreamServerInterceptor()),
		grpc.UnaryInterceptor(auth.UnaryServerInterceptor()),
	)

	return s, nil
}

func (s *Server) Close() (err error) {
	log.Println("initializing gRPC-server shutdown...")
	s.grpc.GracefulStop()
	s.l.Close()
	s.wg.Wait()
	log.Println("gRPC-server has succesfully stoped!")
	return
}

func (s *Server) RunAsync() {
	log.Printf("gRPC-server running on %v", s.cfg.Bind)
	reflection.Register(s.grpc)
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		if err := s.grpc.Serve(s.l); err != nil && err != grpc.ErrServerStopped {
			return
		}
	}()
}

func Wait() {
	var sig = make(chan os.Signal, 2)
	signal.Notify(sig, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)
	<-sig
}

func (s *Server) Grpc() *grpc.Server {
	return s.grpc
}
