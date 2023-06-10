up:
	docker compose --env-file $(env_file) up -d

build:
	go mod tidy
	go build -o $(output) cmd/app/main.go
	$(output)

test:
	go test -v ./...