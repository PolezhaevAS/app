ACCESS_DB = postgresql://postgres:qwerty@127.0.0.1:5432/postgres?sslmode=disable

static:
	staticcheck ./...

vet:
	go vet ./...

docker-compose:
	docker compose -f docker-compose.yaml up --build 

migrate-up:
	migrate -path access/migrates -verbose -database=${ACCESS_DB} up

migrate-down:
	migrate -path access/migrates -verbose -database=${ACCESS_DB} down