#!make

-include .env

export

GOPATH=$(shell go env GOPATH)

setup:
	go install golang.org/x/lint/golint@latest
	go install github.com/golang/mock/mockgen@latest
	
rest-api:
	go run src/backend/cmd/rest_api/main.go

mock:
	rm -rf mocks
	$(GOPATH)/bin/mockgen -source=src/backend/domain/repository.go -destination=mocks/domainrepository.go -package=mocks