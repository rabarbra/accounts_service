package user

import "errors"

var ErrNotFound = errors.New("not found")

type User struct {
	Id    int32
	Email string
}

type Service interface {
	AddUser(email string) (int32, error)
	GetUser(id int32) (User, error)
	ListUsers() (map[int32]User, error)
}
