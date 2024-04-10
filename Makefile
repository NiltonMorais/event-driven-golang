#!make
include .env
export $(shell sed 's/=.*//' .env)

build: 
	@go build -o bin/server cmd/api/main.go

server: build up
	@./bin/server

up:
	@docker-compose up -d	

down: 
	@docker-compose down	

test:
	@go test -v ./...

