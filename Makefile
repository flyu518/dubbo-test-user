
GOPATH:=$(shell go env GOPATH)
PWD:=$(shell pwd)

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: run
run:
	@go run ./main.go

.PHONY: build
build:
	@go build -o ./server ./main.go

.PHONY: docker
docker:
	@docker build -t user:latest .
#	@docker tag user:latest xxx
#	@docker push xxx
