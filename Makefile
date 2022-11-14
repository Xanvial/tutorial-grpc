build-client:
	@echo " >> building client binary"
	@go build -v -o app-client client/main.go

build-server:
	@echo " >> building server binary"
	@go build -v -o app-server server/main.go

run-client: build-client
	@./app-client

run-server: build-server
	@./app-server

gen-proto:
	@protoc proto/*.proto --go_out=plugins=grpc:proto -I proto
