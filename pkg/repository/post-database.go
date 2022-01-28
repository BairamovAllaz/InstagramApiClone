package repository

import (
	"Postresql/structs"
	"database/sql"
	"errors"
)

type PostData struct {
	db *sql.DB
}
func NewPostData(db *sql.DB) *PostData { 
	return &PostData{db : db};
}
func(p *PostData)Addpost(post structs.PostStruct)(int,error){ 
	insertstring :=  `
		INSERT INTO posts(userid,image,comment)VALUES($1,$2,$3)
	`
	_,err := p.db.Exec(insertstring,post.User,post.Image.Filename,post.Comment); 

	if err != nil { 
		return 0, err;
	}
	return 0,nil;
}
func(p *PostData)AddLikeToPostRepo(postid string,userid int)(int,error) { 
		checkstring := `
			SELECT 1 FROM posts WHERE id=$1 AND userid=$2
		`
		var id int;
		err := p.db.QueryRow(checkstring,postid,userid).Scan(&id);
		if err != sql.ErrNoRows { 
			return 0,errors.New("you cant like your post");
		}

		updateString := `
			UPDATE posts SET likecount=likecount+1 WHERE id=$1 
		`
		_,err = p.db.Exec(updateString,postid);

		if err != nil  { 
			return 0,err;
		}
		return id,nil;
}

func(p *PostData)DeletePost(postid string ,userid int)(int,error) { 
	checkstring := `
	SELECT 1 FROM posts WHERE id=$1 AND userid=$2
	`
	var id int;
	err := p.db.QueryRow(checkstring,postid,userid).Scan(&id);
	if err == sql.ErrNoRows { 
		return 0,errors.New("you cant delete other post");
	}
	deleteString := `
		DELETE FROM posts WHERE id=$1 
	`
	_,err = p.db.Exec(deleteString,postid);
	if err != nil { 
		return 0, err
	}
	return id,nil;
}
