package services

import (
	"Postresql/pkg/repository"
	"Postresql/structs"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/crypto/bcrypt"
)

const(
	key = "efoihioivihiohv"
)

type Authservicee struct {
	repos repository.Authrization
}

type claims struct { 
	jwt.StandardClaims
	Email string `json:"email"`
}

func NewAuthService(repo repository.Authrization) *Authservicee { 
	return &Authservicee{repos: repo};
}
func(a *Authservicee) Signin(user structs.User)(int,error) { 
	newpassword, err := HashPassword(user.Passsword);
	
	if err != nil { 
		log.Fatalf("error: %s", err.Error())
	}

	user.Passsword = newpassword;
	fmt.Println("New password: ", user.Passsword);
	return a.repos.Signin(user);
}


func(a *Authservicee)SignUp(user structs.SignUpuser)(string,error) { 
	hashpassword,err := HashPassword(user.Password);
	if err != nil { 
		return "", err;
	} 
	user.Password= hashpassword;
	newuser,err := a.repos.SignUp(user);
	
	if err != nil { 
		return  "", err;
	}

	//*create token for jwt 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,&claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(12 * time.Hour)),
			IssuedAt: jwt.At(time.Now()),
		},
		Email: newuser.Email,
	})

	return token.SignedString([]byte(key));
}


///*Parsing jwt token

func(a *Authservicee)Parsetoken(token string) (string,error) { 
	mytoken,err := jwt.ParseWithClaims(token,&claims{},func(t *jwt.Token) (interface{}, error) {
		if _,ok := t.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil,errors.New("Eroorparsing")
		}
		return []byte(key),nil;

	})
	if err != nil { 
		return "",err;
	}
	claim ,ok := mytoken.Claims.(*claims)

	if !ok { 
		return "",errors.New("errrrr"); 
	}
	return claim.Email,nil;
}




func HashPassword(userpassword string) (string,error) { 
	bytes,err := bcrypt.GenerateFromPassword([]byte(userpassword),14);
	return string(bytes),err;
}
