package core

import (
	"github.com/rabarbra/user"
)

type service struct {
	m map[int32]user.User
}

func NewService() user.Service {
	return &service{
		m: map[int32]user.User{
			1: {Id: 1, Name: "Poly"},
			2: {Id: 2, Name: "SDdsf"},
		},
	}
}
