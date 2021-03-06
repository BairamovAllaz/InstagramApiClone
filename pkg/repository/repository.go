package repository

import (
	"Postresql/structs"
	"database/sql"

)

type Authrization interface {
	Signin(user structs.User)(int,error)
	SignUp(user structs.SignUpuser)(int,error)
}
type Postrepo interface { 
	Addpost(post structs.PostStruct)(int,error)
	AddLikeToPostRepo(postid string,userid int)(int,error)
	DeletePost(postid string ,userid int)(int,error)
	AddDislikeToPost(postid string, userid int)(int,error)
}
///interface struct
type Repository struct {
	Authrization //*interface
	Postrepo
}
func NewRepository(db *sql.DB) *Repository { 
	return &Repository{
		Authrization: NewAuthSql(db),
		Postrepo: NewPostData(db),
	}
}	