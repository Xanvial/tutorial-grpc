# tutorial-grpc

### Run Project
download all necessary dependency with
```
go mod vendor
```

To generate the proto file, use 
```
make gen-proto
```
or
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=paths=source_relative:. proto/*.proto
```

On the root directory, run this command for server:
```
make run-server
```
on other terminal, run client with:
```
make run-client
```

If `make` is not installed, the application can be run using these command:
```
go run server/main.go
go run client/main.go
```

## Directory structure
```
├── client                              # Contains the executable programs for client
│  └── main.go                          # Main program of client app
├── example                             # Example folder
├── go.mod                              # Go module file (collection of Go packages)
├── go.sum                              # Go sum file
├── LICENSE                             # License file
├── Makefile                            # Makefile file
├── proto                               # Collection of proto files, both `.proto` and `.go` format
├── README.md                           # Readme file
└── server
   ├── main.go                          # Main program of server app
   ├── model                            # Collection of struct definition
   └── usecase                          # Main usecase of server app
```