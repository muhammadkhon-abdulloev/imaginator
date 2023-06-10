IMAGINATOR
---
### Used libs:
* [go-grpc](https://google.golang.org/grpc) - the Go implementation of gRPC
* [uber/fx](https://github.com/uber-go/fx) - dependency injection system for Go.
* [go-grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware) - [go-grpc](https://google.golang.org/grpc) middlewares: interceptors, helpers and utilities.
* [go-envconfig](https://github.com/sethvargo/go-envconfig) - Envconfig populates struct field values based on environment variables or arbitrary lookup functions.


Build and run:
```shell
make build output=WHERE_OUTPUT
```

Using docker:
```shell
make up env_file=PATH_TO_ENV_FILE
```

Run tests:
```shell
make test
```

Generate .proto files
```shell
make protoc
```
