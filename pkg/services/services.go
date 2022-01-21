package services

import (
	"Postresql/pkg/repository"
	"Postresql/structs"
)

type Authservice interface {
	Signin(user structs.User)(int,error)
	SignUp(user structs.SignUpuser)(string,error)
	Parsetoken(token string) (string,error)
}

type Service struct {
	Authservice
}

func NewService(repos *repository.Repository) *Service { 
	return &Service{
		Authservice: NewAuthService(repos.Authrization),
	}
}
