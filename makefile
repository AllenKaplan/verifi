.DEFAULT_GOAL := help

help:
	@ echo "	go-build: compile go binary"
	@ echo "	build: build container with deps"
	@ echo "	up: start docker containers"
	@ echo "	setup: run a setup/configuration script that will prepopulate the database"
	@ echo "	update: restart app container with new code changes"
	@ echo "	down: kill all containers"
	@ echo "	clean: DO NOT USE UNLESS YOU KNOW WHAT YOU ARE DOING"
	@ echo "	       This command will stop all running containers, and delete every single"
	@ echo "	       container image from your system"
	@ echo "	shell: interactive shell for the go container, make sure the container is running"
	@ echo "	mock-profile: generate mock-gen test for profile service"

build-go:
	go build -o app ./cmd

run: build-go
	./app

build-migrate:
	go build -o migrate ./cmd/migrate

#build:
#	DOCKER_BUILDKIT=1 docker build -t 'radar_backend' .
#
#start:
#	docker-compose up -d #> /dev/null
#
#up: build start
#
#down:
#	docker-compose stop
#	docker-compose down
#
#clean: down
#	docker kill $(docker ps -q) || echo "return 0"
#	docker rm $(docker ps -a -q) || echo "return 0"
#	docker volume rm $(docker volume ls -q) || echo "return 0"
#	docker system prune
#
#protoc:
#	protoc --proto_path=./protos --go_out=./protos --go_opt=paths=source_relative --go-grpc_out=./protos --go-grpc_opt=paths=source_relative --grpc-gateway_out ./protos --grpc-gateway_opt paths=source_relative ./protos/*.proto
#
#mocks: mock-profile mock-match mock-recommendation mock-messaging mock-auth
#
#mock-reporting:
#	 mockgen -source=./reporting/repo/db.go -destination=./reporting/mocks/mock_db.go -package=mocks

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

lint: fmt vet