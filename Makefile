ACCESS_DB = postgresql://postgres:qwerty@127.0.0.1:5432/access?sslmode=disable
AUTH_DB = postgresql://postgres:qwerty@127.0.0.1:5432/users?sslmode=disable

static:
	staticcheck ./...

vet:
	go vet ./...

docker-compose:
	docker compose -f docker-compose.yaml up --remove-orphans
	

docker-services:
	docker compose -f docker-compose.services.yaml up 

migrate-up:
	migrate -path access/migrates -verbose -database=${ACCESS_DB} up
	migrate -path auth/migrates -verbose -database=${AUTH_DB} up

migrate-down:
	migrate -path access/migrates -verbose -database=${ACCESS_DB} down
	migrate -path auth/migrates -verbose -database=${AUTH_DB} down