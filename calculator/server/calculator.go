package main

import (
	"context"
	"fmt"
	pb "github.com/wfabjanczuk/grpcGolang/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"math"
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

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max invoked")

	var count, max int64

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Fatalf("Error while receiving client stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)
		if count == 0 {
			max = req.Number
			sendMaxResponse(max, stream)
		} else if req.Number > max {
			max = req.Number
			sendMaxResponse(max, stream)
		}

		count++
	}

	return nil
}

func sendMaxResponse(max int64, stream pb.CalculatorService_MaxServer) {
	log.Printf("New maximum: %d\n", max)
	err := stream.Send(&pb.MaxResponse{
		Max: max,
	})
	if err != nil {
		log.Fatalf("Error while sending data to client: %v\n", err)
	}
}

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt invoked with %v\n", in)

	if in.Number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", in.Number),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(in.Number)),
	}, nil
}
