package main

import (
	"context"
	pb "github.com/wfabjanczuk/grpcGolang/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
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

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg invoked")
	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling doAvg: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving Avg: %s\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax invoked")
	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while calling doMax: %v\n", err)
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
				break
			}
			log.Printf("Received: %v\n", res)
		}
		close(waitc)
	}()

	<-waitc
}

func doSqrt(n int64, c pb.CalculatorServiceClient) {
	log.Println("doSqrt invoked")
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n,
	})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %d\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("Invalid argument")
			}
			return
		} else {
			log.Fatalf("Non gRPC error: %v\n", err)
		}
	}

	log.Printf("Result: %f\n", res.Result)
}
