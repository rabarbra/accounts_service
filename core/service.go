package core

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rabarbra/accounts_service/logg"
	pb "github.com/rabarbra/accounts_service/proto"
)

var ErrNotFound = errors.New("not found")

type User struct {
	Id    int32
	Email string
}

type UserService interface {
	AddUser(email string) (int32, error)
	GetUser(id int32) (User, error)
	ListUsers() (map[int32]User, error)
}

type Service struct {
	pb.UserServiceServer
	logg *logg.Logger
	db   *pgxpool.Pool
}

func NewService() *Service {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		log.Fatalf("Unable to connection to database: %v\n", err)
	}
	defer pool.Close()
	log.Println("Connected!")
	logger := logg.New()

	return &Service{
		logg: logger,
		db:   pool,
	}
}

func (s *Service) ListUsers(ctx context.Context, _ *pb.ListRequest) (result User, err error) {
	var email string
	s.db.QueryRow(context.Background(), "select email from users where id = $1", id).Scan(&email)
	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
	}
	user := User{
		Id:    id,
		Email: email,
	}
	return user, err
}

func (s *Service) AddUser(ctx context.Context, req *pb.AddRequest) (id int32, err error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	_ = tx.QueryRow(ctx, "INSERT INTO users email VALUES $1 RETURNING id", req.Email).Scan(&id)

	if err = tx.Commit(ctx); err != nil {
		return 0, err
	}

	return &pb.Reply{}, nil
}

//func (s *Service) ListUsers() (list map[int32]models.User, err error) {
//	return s.m, nil
//}
