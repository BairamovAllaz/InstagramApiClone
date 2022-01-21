package structs

type User struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Passsword string `json:"password"`
}

type SignUpuser struct { 
	Email string `json:"email"`
	Password string `json:"password"`
}
