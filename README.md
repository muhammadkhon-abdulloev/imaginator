IMAGINATOR

To start:
```shell
go mod tidy
go run cmd/app/main.go
```

OR using docker

```shell
docker compose --env-file .env.example up -d
```

Run tests
```shell
go test -v ./...
```