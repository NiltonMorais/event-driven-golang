#!make
include .env
export $(shell sed 's/=.*//' .env)

build: 
	@go build -o bin/server cmd/api/main.go
	@go build -o bin/consumer cmd/consumer/main.go

server: build
	@./bin/server
	
consumer: build
	@./bin/consumer

test:
	@go test -v ./...

up:
	@docker-compose up -d	

