package main

import (
	"context"
	pb "github.com/wfabjanczuk/grpcGolang/blog/proto"
	"log"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Printf("deleteBlog was invoked with %v\n", id)

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error while deleting: %v\n", err)
	}

	log.Printf("Blog was deleted: %s", id)
}
