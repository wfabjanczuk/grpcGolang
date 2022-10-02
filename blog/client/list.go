package main

import (
	"context"
	pb "github.com/wfabjanczuk/grpcGolang/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Println("readBlog was invoked")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Println(res)
	}
}
