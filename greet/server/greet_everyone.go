package main

import (
	"fmt"
	pb "github.com/wfabjanczuk/grpcGolang/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)
		res := fmt.Sprintf("Hello %s!", req.FirstName)
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
