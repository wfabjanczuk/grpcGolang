package main

import (
	"context"
	pb "github.com/wfabjanczuk/grpcGolang/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet invoked")
	reqs := []*pb.GreetRequest{
		{FirstName: "Wojtek"},
		{FirstName: "Wojciech"},
		{FirstName: "Test"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
