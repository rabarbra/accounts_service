package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	pb "github.com/rabarbra/account_service/proto"
)

const (
	port string = "8080:"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.UnimplementedUserServiceServer(s, $server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}