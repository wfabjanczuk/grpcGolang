package main

import (
	"context"
	pb "github.com/wfabjanczuk/grpcGolang/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Wojtek",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
