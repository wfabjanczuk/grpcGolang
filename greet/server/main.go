package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/wfabjanczuk/grpcGolang/greet/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v\n", addr, err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
