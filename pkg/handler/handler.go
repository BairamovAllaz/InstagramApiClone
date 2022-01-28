package handler

import (
	"Postresql/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services services.Service) *Handler {
	return &Handler{services :&services}; 
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/user")
	{
		v1.POST("/signin",h.Signin)
		v1.POST("/signup",h.Signup)
		v1.GET("/logout",h.Logout)
	}

	v2 := router.Group("/add",h.userIdenity) 
	{ 
		v2.POST("/post",h.Addpost)
	}

	v3 := router.Group("/do",h.userIdenity) 
	{ 
		v3.POST("/like/:id",h.AddLikeToPost)
		v3.DELETE("/delete/:id",h.DeletPost)
	}

	router.StaticFS("/file", http.Dir("Images"))
	return router
}
