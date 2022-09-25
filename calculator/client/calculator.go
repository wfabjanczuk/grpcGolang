package main

import (
	"context"
	pb "github.com/wfabjanczuk/grpcGolang/calculator/proto"
	"io"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Result: %d\n", res.Result)
}

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes invoked")

	req := &pb.PrimesRequest{
		Number: 210,
	}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Primes: %d\n", msg.Factor)
	}
}
