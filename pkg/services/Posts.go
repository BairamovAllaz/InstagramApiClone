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

func (p *Posts) Addpost(post structs.PostStruct,c *gin.Context) (string, error) {
	// filepath := SaveImageToFolder(post.Image)
	filename := post.Image.Filename

	// out, err := os.Create("Images/" + filename);
	// if err != nil {
	// 	log.Fatalf("error while saving file: %s", err.Error())
	// 	return "",nil;
	// }
	// defer out.Close()
	err := c.SaveUploadedFile(post.Image,"Images/" + filename);
	if err != nil { 
		log.Fatalf("error %s", err.Error());
	}
	filepath := "http://localhost:8000/file/" + filename

	return filepath, nil
}

// func SaveImageToFolder(image *multipart.FileHeader) string {
// 	filename := image.Filename;

// 	out,err := os.Create("hello.txt");
// 	if err != nil {
// 		log.Fatalf("error while saving file: %s", err.Error())
// 		os.Exit(1);
// 	}

// 	defer out.Close();

// 	filepath := "http://localhost:8000/file/" + filename;

// 	return filepath;
// }
