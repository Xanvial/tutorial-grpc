build-client:
	@go build -v -o example-client client/main.go

build-server:
	@go build -v -o example-server server/main.go

run-client: build-client
	@./example-client

run-server: build-server
	@./example-server

gen-proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=paths=source_relative:. proto/*.proto
