# folder structure

```
.
├── README.md
├── client
│   └── client.go
├── go.mod
├── go.sum
├── protobuf
│   ├── status.pb.go
│   ├── status.proto
│   └── status_grpc.pb.go
└── server
    └── server.go
```

## protobuf

contains `.proto` file and generated `.pb.go` & `grpc.pb.go` file

#### command to generate
```
protoc --go_out=. --go_opt=paths=source_relative \                                                                                    master
    --go-grpc_out=. --go-grpc_opt=paths=source_relative status.proto
```

## server

contains code for grpc server

## client

contains code for grpc client

## References

https://grpc.io/docs/languages/go/basics/

https://github.com/grpc/grpc-go/tree/master/examples

