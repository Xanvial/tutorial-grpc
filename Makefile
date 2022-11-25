build-client:
	@go build -v -o app-client client/main.go

build-server:
	@go build -v -o app-server server/main.go

run-client: build-client
	@./app-client

run-server: build-server
	@./app-server

gen-proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=paths=source_relative:. proto/*.proto

gen-proto-old:
	@protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative proto/*.proto
