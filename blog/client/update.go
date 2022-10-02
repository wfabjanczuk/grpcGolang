package main

import (
	"context"
	pb "github.com/wfabjanczuk/grpcGolang/blog/proto"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Printf("updateBlog was invoked with %v\n", id)

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Wojtek",
		Title:    "A new title",
		Content:  "A new content",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated!")
}
