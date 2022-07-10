package main

import (
	"app/access/pkg/access"
	pb_access "app/access/pkg/access/gen"
	"app/auth/internal/config"
	db "app/auth/internal/database"
	server_auth "app/auth/internal/server"
	service_auth "app/auth/internal/service"
	pb "app/auth/pkg/proto/gen"
	"app/internal/broker"
	grpc_server "app/internal/server"
	grpc_auth "app/internal/server/auth"
	"app/internal/token"
	"log"
)

func main() {
	var (
		cfg      = config.New().Load()
		err      error
		jwt      *token.Source
		database db.Repository
		grpc     *grpc_server.Server
		grpcAuth *grpc_auth.Auth
		server   *server_auth.Server
		service  *service_auth.Auth
	)

	if database, err = db.New(cfg.DB); err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	serviceDesc := server_auth.Rules(pb.Auth_ServiceDesc)

	jwt, err = token.New(cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	grpcAuth = grpc_auth.New(jwt, serviceDesc)
	if err != nil {
		log.Fatal(err)
	}

	broker, err := broker.New(pb_access.AccessBroker_ServiceDesc, cfg.Broker)
	if err != nil {
		log.Fatal(err)
	}

	grpc, err = grpc_server.New(cfg.GRPC, grpcAuth)
	if err != nil {
		log.Fatal(err)
	}

	accessClient := access.New(broker)

	service = service_auth.New(database, cfg, jwt, accessClient)
	server = server_auth.New(service)

	pb.RegisterAuthServer(grpc.Grpc(), server)
	grpc.RunAsync()
	defer grpc.Close()

	grpc_server.Wait()
}
