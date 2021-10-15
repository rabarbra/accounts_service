package core

import (
	models "github.com/rabarbra/accounts_service/models"
	pb "github.com/rabarbra/accounts_service/proto"
)

type Service struct {
	pb.UserServiceServer
	logg *logg.Logger
	db   controller.UserController
}

func NewService() *Service {
	return &Service{
		logg: logg,
		db:   controller.New(db),
	}

}

func (s *service) GetUser(id int32) (result models.User, err error) {
	if result, ok := s.m[id]; ok {
		return result, nil
	}
	return result, models.ErrNotFound
}

func (s *service) AddUser(email string) (id int32, err error) {
	s.m[s.next_id] = models.User{Id: s.next_id, Email: email}
	s.next_id += 1
	return s.next_id - 1, nil
}

func (s *service) ListUsers() (list map[int32]models.User, err error) {
	return s.m, nil
}
