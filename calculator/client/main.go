package main

import (
	pb "github.com/wfabjanczuk/grpcGolang/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to %s: %v\n", addr, err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	//doSum(c)
	//doPrimes(c)
	//doAvg(c)
	//doMax(c)
	doSqrt(10, c)
	doSqrt(-1, c)
}
