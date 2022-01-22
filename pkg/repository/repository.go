package repository

import (
	"Postresql/structs"
	"database/sql"

)

type Authrization interface {
	Signin(user structs.User)(int,error)
	SignUp(user structs.SignUpuser)(int,error)
}

///interface struct
type Repository struct {
	Authrization //*interface
}

func NewRepository(db *sql.DB) *Repository { 
	return &Repository{
		Authrization: NewAuthSql(db),
	}
}
