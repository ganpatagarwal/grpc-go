package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	pb "github.com/ganpatagarwal/grpc-go/status"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var (
	filename = flag.String("filename", "records", "File to save records")
	port     = flag.Int("port", 10000, "The server port")
)

type statusServiceServer struct {
	pb.UnimplementedStatusServiceServer
	recordFileName string
	statusRecords  *pb.StatusReport
}

func (s *statusServiceServer) ReadStatus(ctx context.Context, query *pb.StatusQuery) (*pb.Status, error) {
	// Read the existing records.
	in, err := ioutil.ReadFile(s.recordFileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("%s: File not found.  Creating new file.\n", s.recordFileName)
		} else {
			return nil, fmt.Errorf("error in reading records file")
		}
	}

	// unmarshal records
	if err := proto.Unmarshal(in, s.statusRecords); err != nil {
		return nil, err
	}

	for _, st := range s.statusRecords.Status {
		if st.Fqdn == query.Fqdn {
			return st, nil
		}
	}

	return nil, fmt.Errorf("not found")
}

func (s *statusServiceServer) UpdateStatus(ctx context.Context, st *pb.Status) (*pb.StatusUpdate, error) {
	// Read the existing records.
	in, err := ioutil.ReadFile(s.recordFileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("%s: File not found.  Creating new file.\n", s.recordFileName)
		} else {
			return &pb.StatusUpdate{Status: false}, fmt.Errorf("error in reading records file")
		}
	}

	// unmarshal records
	if err := proto.Unmarshal(in, s.statusRecords); err != nil {
		return &pb.StatusUpdate{Status: false}, err
	}

	s.statusRecords.Status = append(s.statusRecords.Status, st)

	// Write the new records back to disk.
	out, err := proto.Marshal(s.statusRecords)
	if err != nil {
		return &pb.StatusUpdate{Status: false}, err
	}
	if err := ioutil.WriteFile(s.recordFileName, out, 0644); err != nil {
		return &pb.StatusUpdate{Status: false}, err
	}

	return &pb.StatusUpdate{Status: true}, nil
}

func newStatusServiceServer(filename string) *statusServiceServer {
	return &statusServiceServer{
		recordFileName: filename,
		statusRecords:  &pb.StatusReport{},
	}
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStatusServiceServer(grpcServer, newStatusServiceServer(*filename))
	log.Println("Starting GRPC server on port: ", *port)
	grpcServer.Serve(lis)
}
