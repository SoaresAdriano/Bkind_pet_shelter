#!make

-include .env

export

GOPATH=$(shell go env GOPATH)

setup:
	go install golang.org/x/lint/golint@latest
	go install github.com/golang/mock/mockgen@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	
grpc-api_petCadaster_service:
	go run src/services/petCadaster_service/cmd/api/main.go

mock:
	rm -rf mocks
	$(GOPATH)/bin/mockgen -source=src/services/domain/repository.go -destination=mocks/domainrepository.go -package=mocks

up_back_db:
	cd src/db/mysql && docker compose up -d 

shutdown_back_db:
	cd src/db/mysql && docker compose down 

proto_grpc_petCadaster_service:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	src/services/petCadaster_service/ports/grpc/protos/*.proto
    
