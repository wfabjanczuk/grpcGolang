package main

import (
	"context"
	pb "github.com/wfabjanczuk/grpcGolang/calculator/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
