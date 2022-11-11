build-client:
	@echo " >> building client binary"
	@go build -v -o example-client client/main.go

build-server:
	@echo " >> building server binary"
	@go build -v -o example-server server/main.go

run-client: build-client
	@./example-client

run-server: build-server
	@./example-server
