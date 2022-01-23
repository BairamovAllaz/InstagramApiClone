package services

import (
	"Postresql/pkg/repository"
	"Postresql/structs"
)

type Posts struct {
	repo repository.Postrepo
}

func NewPosts(repo repository.Postrepo) *Posts { 
	return &Posts{repo: repo}
}

func(p *Posts)Addpost(post structs.PostStruct)(int,error){ 
	return 0,nil;
}