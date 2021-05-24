package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/ganpatagarwal/grpc-go/protobuf"

	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error in connecting to server: %s", err)
	}
	defer conn.Close()

	// create client
	client := pb.NewStatusServiceClient(conn)
	ctx := context.Background()

	// write data to server
	_, err = client.UpdateStatus(ctx,
		&pb.Status{
			Fqdn:   "localhost",
			Ip:     "127.0.0.1",
			Status: true,
		})

	// read data from server
	st, err := client.ReadStatus(ctx,
		&pb.StatusQuery{Fqdn: "localhost"})

	if err != nil {
		log.Println(err)
	} else {
		log.Println(st)
	}
}
