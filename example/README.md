# Example for tutorial-grpc
Example program that implements server and client of grpc, with basic interceptor on server side

### Run Project
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
├── Makefile                            # Makefile file
├── proto                               # Collection of proto files, both `.proto` and `.go` format
├── README.md                           # Readme file
└── server                              # Collection of files to be stored on the server
   ├── hello                            # Server logic classes
   ├── interceptor                      # Collection of interceptor
   └── main.go                          # Main program of server app
```

### Libraries
Main libraries:
- gRPC  - https://github.com/grpc/grpc-go
        - https://github.com/golang/protobuf
        - https://github.com/grpc-ecosystem/go-grpc-middleware
