package main

import (
	"app/access/internal/config"
	db "app/access/internal/database"
	server_access "app/access/internal/server"
	service_access "app/access/internal/service"
	pb "app/access/pkg/proto/gen"
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

	serviceDesc := server_access.Rules(&pb.Access_ServiceDesc)

	jwt, err = token.New(cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	grpc_auth = auth.New(jwt, serviceDesc, "admin")

	grpc, err = grpc_server.New(cfg.GRPC, grpc_auth)
	if err != nil {
		log.Fatal(err)
	}

	service = service_access.New(database)
	server = server_access.New(service)

	pb.RegisterAccessServer(grpc.Grpc(), server)
	grpc.RunAsync()
	defer grpc.Close()

	grpc_server.Wait()
}
