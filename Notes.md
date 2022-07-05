# Ports and Adapters architecture 
Here I will describe the insights and some value tips to use when applying the ports and adapters architecture.

## Development order

1. Domain
2. Application
3. Ports
4. Adapters

## Resources

[Ports & Adapters architecture on example](https://wkrzywiec.medium.com/ports-adapters-architecture-on-example-19cab9e93be7)

## Working with protoc-gen-go

I've had some problems with protoc-gen-go commands, so, let me do a resume that you must do to it works ok:

1. go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
2. go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
3. Put `export PATH="$PATH:$(go env GOPATH)/bin"` on zshrc or bashrc file
4. go mod tidy
5. Use command `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative <your proto path>/<.proto file>`