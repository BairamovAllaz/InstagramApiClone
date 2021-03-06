package services

import (
	"Postresql/pkg/repository"
	"Postresql/structs"

	"github.com/gin-gonic/gin"
)

type Authservice interface {
	Signin(user structs.User)(int,error)
	SignUp(user structs.SignUpuser)(string,error)
	Parsetoken(token string) (int,error)
}


type Postservice interface { 
	Addpost(post structs.PostStruct,c *gin.Context,userid int)(string,error)
	AddLikeToPostService(postid string,userid int)(int,error)
	DeletPost(postid string,userid int) (int,error)
	AddDislikeToPost(postid string, userid int)(int,error)
}

type Subscribeservice interface{ 
	SubscribeTo(id string,userid int)(int,error)
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
