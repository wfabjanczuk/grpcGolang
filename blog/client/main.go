package main

import (
	pb "github.com/wfabjanczuk/grpcGolang/blog/proto"
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

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)
	//readBlog(c, "non-existing-id")
	updateBlog(c, id)
	//updateBlog(c, "non-existing-id")
	listBlogs(c)
	deleteBlog(c, id)
}
