package main

import (
	"context"
	pb "github.com/wfabjanczuk/grpcGolang/calculator/proto"
	"log"
)

func doCalculate(c pb.CalculatorServiceClient) {
	log.Printf("doCalculate invoked")
	res, err := c.Calculate(context.Background(), &pb.CalculatorRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	})
	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}

	log.Printf("Result: %d\n", res.Result)
}
