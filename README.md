IMAGINATOR

Для билда и запуска:
```shell
make build output=WHERE_OUTPUT
```

Для запуска в docker

```shell
make up env_file=PATH_TO_ENV_FILE
```

Для запуска тестов
```shell
make test
```

Для генерации .proto файла
```shell
protoc -I . --go_out=. --go-grpc_out=. proto/imaginator.proto
```