package services

import (
	"Postresql/pkg/repository"
	"Postresql/structs"
)

type Authservice interface {
	Signin(user structs.User)(int,error)
	SignUp(user structs.SignUpuser)(string,error)
	Parsetoken(token string) (int,error)
}

type Postservice interface { 
	Addpost(post structs.PostStruct)(int,error)
}

type Service struct {
	Authservice
	Postservice
}

func NewService(repos *repository.Repository) *Service { 
	return &Service{
		Authservice: NewAuthService(repos.Authrization),
		Postservice: NewPosts(repos.Postrepo),
	}
}
