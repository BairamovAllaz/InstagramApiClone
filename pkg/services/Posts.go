package services

import (
	"Postresql/pkg/repository"
	"Postresql/structs"
	"log"

	"github.com/gin-gonic/gin"
	// "fmt"
	// "io"
)
type Posts struct {
	repo repository.Postrepo
}

func NewPosts(repo repository.Postrepo) *Posts {
	return &Posts{repo: repo}
}

func (p *Posts) Addpost(post structs.PostStruct,c *gin.Context,userid int) (string, error) {
	filename := post.Image.Filename
	err := c.SaveUploadedFile(post.Image,"Images/" + filename);
	if err != nil { 
		log.Fatalf("error %s", err.Error());
	}
	filepath := "http://localhost:8000/file/" + filename
	post.User = userid;
	_,err = p.repo.Addpost(post);
	if err != nil { 
		return "",err;
	}
	return filepath, nil
}

func(p *Posts)AddLikeToPostService(postid string,userid int)(int,error) { 
	return p.repo.AddLikeToPostRepo(postid,userid);
}

func(p *Posts)DeletPost(postid string,userid int) (int,error) { 
	return p.repo.DeletePost(postid,userid);
}


func(p *Posts)AddDislikeToPost(postid string, userid int)(int,error){  

	return p.repo.AddDislikeToPost(postid,userid);

}