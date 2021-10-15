package core

import (
	"github.com/rabarbra/models"
)

type service struct {
	m       map[int32]models.User
	next_id int32
}

func NewService() models.Service {
	return &service{
		m: map[int32]models.User{
			1: {Id: 1, Email: "Poly@fv.tre"},
			2: {Id: 2, Email: "SDdsf@ret.tr"},
		},
		next_id: 3,
	}
}

func (s *service) GetUser(id int32) (result models.User, err error) {
	if result, ok := s.m[id]; ok {
		return result, nil
	}
	return result, models.ErrNotFound
}

func (s *service) AddUser(email string) (id int32, err error) {
	s.m[s.next_id] = models.User{Id: s.next_id, email: email}
	s.next_id += 1
	return s.next_id - 1, nil
}

func (s *service) ListUsers() (list map[int32]models.User, err error) {
	return s.m, nil
}
