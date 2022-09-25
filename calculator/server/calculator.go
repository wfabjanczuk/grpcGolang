package main

import (
	"context"
	pb "github.com/wfabjanczuk/grpcGolang/calculator/proto"
	"io"
	"log"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes invoked with %v\n", in)

	var k, n int64
	k = 2
	n = in.Number

	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimesResponse{
				Factor: k,
			})
			n = n / k
		} else {
			k++
		}
	}

	return nil
}

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg invoked")

	var sum, n int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(n),
			})
		} else if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		log.Printf("Receiving: %v\n", req)
		sum += req.Number
		n++
	}

	return nil
}
