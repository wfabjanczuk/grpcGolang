package main

import (
	"context"
	pb "github.com/wfabjanczuk/grpcGolang/blog/proto"
	"log"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Printf("readBlog was invoked with %v\n", id)

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("Error happened while reading: %v\n", err)
		return nil
	}

	log.Printf("Blog was read: %s", res)
	return res
}
