# tutorial-grpc

### Run Project
download all necessary dependency with
```
go mod vendor
```

On the root directory, run this command for server:
```
make run-server
```
on other terminal, run client with:
```
make run-client
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