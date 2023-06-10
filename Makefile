up:
	docker compose --env-file $(env_file) up -d --build

build:
	go mod tidy
	go build -o $(output) cmd/app/main.go
	$(output)

test:
	go test -v ./...

protoc:
	protoc -I . --go_out=. --go-grpc_out=. proto/imaginator.proto