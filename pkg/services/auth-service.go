package services

import (
	"Postresql/pkg/repository"
	"Postresql/structs"
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

const(
	key = "efoihioivihiohv"
)

type Authservicee struct {
	repos repository.Authrization
}

type claims struct { 
	jwt.StandardClaims
	Id int `json:"id"`
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
	fmt.Println("User password ", user.Password)
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
		Id: newuser,
	})

	return token.SignedString([]byte(key));
}


///*Parsing jwt token

func(a *Authservicee)Parsetoken(token string) (int,error) { 
	mytoken,err := jwt.ParseWithClaims(token,&claims{},func(t *jwt.Token) (interface{}, error) {
		if _,ok := t.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil,errors.New("Eroorparsing")
		}
		return []byte(key),nil;

	})
	if err != nil { 
		return 0,err;
	}
	claim ,ok := mytoken.Claims.(*claims)

	if !ok { 
		return 0,errors.New("errrrr"); 
	}
	return claim.Id,nil;
}




func HashPassword(userpassword string) (string,error) { 
	hash := sha1.New()
	hash.Write([]byte(userpassword))
	return fmt.Sprintf("%x", hash.Sum([]byte("hfiuewfiuwuie"))),nil
}
