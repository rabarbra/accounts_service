package user_int

import (
	"net"
    "golang.org/x/net/context"

    "google.golang.org/grpc"
	pb "github.com/rabarbra/accounts_service/proto"
)

const (
	port: ":50051"
)

type server struct {
	pb.UserServiceServer
}
