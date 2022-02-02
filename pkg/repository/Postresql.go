package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DBname   string
}
func NewPostSql(c Config)(*sql.DB,error) { 
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    c.Host, c.Port, c.Username, c.Password, c.DBname)

	db, err := sql.Open("postgres", psqlInfo)
    if err != nil { 
        return nil,err;
    }
    if err := db.Ping();err != nil {
        fmt.Println("error: ", err.Error());
    }
    return db,nil;
}
