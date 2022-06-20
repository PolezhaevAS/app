static:
	staticcheck ./...

vet:
	go vet ./...

docker-compose:
	docker compose -f docker-compose.yaml up --build 