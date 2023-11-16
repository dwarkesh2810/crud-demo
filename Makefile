postgres:
	@docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:16-alpine

start:
	@docker start postgres16

stop:
	@docker stop postgres16

createdb:
	@docker exec -it postgres16 createdb --username=root --owner=root blog

dropdb:
	@docker exec -it postgres16 dropdb blog
	
clear:
	@docker system prune -a

build:
	@go build -o bin/api

run: build
	@./bin/api

test:
	@go test -v ./...
