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
│   └── status_grpc.pb.go
├── server
│   └── server.go
└── status
    ├── status.pb.go
    ├── status.proto
    └── status_grpc.pb.go
```

## status

contains `.proto` file and generated `.pb.go` & `grpc.pb.go` file

## server

contains code for grpc server

## client

contains code for grpc client

## References

https://grpc.io/docs/languages/go/basics/

https://github.com/grpc/grpc-go/tree/master/examples

