package main

import (
	"context"
	pb "github.com/wfabjanczuk/grpcGolang/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Printf("doSum invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Result: %d\n", res.Result)
}
