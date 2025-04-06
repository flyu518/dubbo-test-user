
GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

install-protoc-triple:
	@go install github.com/dubbogo/protoc-gen-go-triple/v3@latest

.PHONY: protoc
protoc:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-triple_out=. --go-triple_opt=paths=source_relative \
    ./api/user.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o ./server ./cmd/server/main.go

.PHONY: docker
docker:
	@docker build -t user:latest .
#	@docker tag user:latest xxx
#	@docker push xxx
