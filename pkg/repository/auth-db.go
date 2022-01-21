package repository

import (
	"Postresql/structs"
	"database/sql"
	"fmt"
)

type Authmongo struct {
	db *sql.DB
}
 
//*create new constructor

func NewAuthSql(db *sql.DB) *Authmongo { 
	return &Authmongo{db: db}
}


func(a *Authmongo) Signin(user structs.User)(int,error) { 
	insertstring := `INSERT INTO users(username,email,password)
	VALUES($1,$2,$3) RETURNING id`

	var id int;
	
	err := a.db.QueryRow(insertstring,user.Username,user.Email,user.Passsword).Scan(&id);

	if err != nil { 
		return 0,err;
	}
	fmt.Println("id: ",id); 
	return id,nil; 
}

func(a *Authmongo)SignUp(user structs.SignUpuser)(structs.User,error) { 
	querystring := `
		SELECT username,email,password FROM users WHERE email=$1 AND password=$2`
	var myuser structs.User;

	err := a.db.QueryRow(querystring,user.Email,user.Password).Scan(&myuser.Username,&myuser.Email,&myuser.Passsword);

	if err != nil {
		return myuser,nil;
	}
	return myuser,nil;
}



