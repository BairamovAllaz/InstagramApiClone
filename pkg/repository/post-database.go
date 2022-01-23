package repository

import (
	"Postresql/structs"
	"database/sql"
)

type PostData struct {
	db *sql.DB
}
func NewPostData(db *sql.DB) *PostData { 
	return &PostData{db : db};
}

func(p *PostData)Addpost(post structs.PostStruct)(int,error){ 
	return 0,nil;
}