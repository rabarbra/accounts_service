package main

import (
	"log"
	"net"

	pb "github.com/rabarbra/accounts_service/proto"
	"google.golang.org/grpc"
)

const port = ":50051"

type server struct {
	pb.UserServiceServer
	//logg *logg.Logger
	//db controller.UserController
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
