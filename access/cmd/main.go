package main

import (
	broker_auth "app/access/internal/broker"
	"app/access/internal/config"
	db "app/access/internal/database"
	server_access "app/access/internal/server"
	service_access "app/access/internal/service"
	broker_pb "app/access/pkg/access/gen"
	pb "app/access/pkg/proto/gen"
	"app/internal/broker"
	grpc_server "app/internal/server"
	auth "app/internal/server/auth"
	token "app/internal/token"
	"context"
	"flag"
	"log"
)

func main() {

	var (
		ctx       = context.Background()
		cfg       = config.New().Load()
		err       error
		jwt       *token.Source
		database  db.Repository
		grpc      *grpc_server.Server
		grpc_auth *auth.Auth
		server    *server_access.Server
		service   *service_access.Access
	)

	init := flag.Bool("init", false, "first start")
	flag.Parse()

	if database, err = db.New(cfg.DB); err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if *init {
		log.Println("Init app")
		err := database.FirstStart(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Start app")

	serviceDesc := server_access.Rules(pb.Access_ServiceDesc)
	broker, err := broker.New(broker_pb.AccessBroker_ServiceDesc, cfg.Broker)
	if err != nil {
		log.Fatal(err)
	}

	jwt, err = token.New(cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	grpc_auth = auth.New(jwt, serviceDesc)

	grpc, err = grpc_server.New(cfg.GRPC, grpc_auth)
	if err != nil {
		log.Fatal(err)
	}

	service = service_access.New(database)
	server = server_access.New(service)
	authBroker := broker_auth.New(broker, service)

	authBroker.Run()
	defer authBroker.Stop()

	pb.RegisterAccessServer(grpc.Grpc(), server)
	grpc.RunAsync()
	defer grpc.Close()

	grpc_server.Wait()
}
