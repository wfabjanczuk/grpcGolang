package main

import (
	"context"
	pb "github.com/wfabjanczuk/grpcGolang/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone invoked")
	reqs := []*pb.GreetRequest{
		{FirstName: "Wojtek"},
		{FirstName: "Wojciech"},
		{FirstName: "Test"},
	}

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Printf("Error while receiving: %v\n", err)
			}
			log.Printf("Received: %v\n", res)
		}
		close(waitc)
	}()

	<-waitc
}
