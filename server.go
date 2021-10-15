package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = ":50051"

type server struct {
	pb
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.registerGreeterServer(s, &server{})
}
