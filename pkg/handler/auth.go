package handler

import (
	"Postresql/structs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Signin(c *gin.Context) {
	var user structs.User;

	if err := c.BindJSON(&user);err != nil { 
		log.Fatalf("error: %s", err.Error()); 
	}
	
	id,err := h.services.Signin(user);

	if err != nil { 
		log.Fatalf("error: %s", err.Error()); 
	}

	c.JSON(http.StatusOK,gin.H{ 
		"id" : id,
	})
}
func (h *Handler) Signup(c *gin.Context) {
	var user structs.SignUpuser;

	if err:=c.BindJSON(&user);err != nil { 
		log.Fatalf("error %s", err.Error()); 
	}


	token,err := h.services.SignUp(user); 

	if err != nil { 
		log.Fatalf("error : %s", err.Error()); 
	}

	c.JSON(http.StatusOK,gin.H{ 
		"result" : token,
	})

}
func (h *Handler) Logout(c *gin.Context) {

}
func (h *Handler) ResetPassword(c *gin.Context) {

}

func(h *Handler)Home(c *gin.Context) { 
	email, ok := c.Get("Userid")

	if !ok {
		log.Println(ok)
	}
	idInt, ok := email.(int)
	if !ok {
		log.Println(ok)
	}

	c.JSON(http.StatusOK,gin.H { 
		"email" : idInt,
	})


}

