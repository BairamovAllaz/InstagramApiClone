package handler

import (
	"Postresql/pkg/services"

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
	router.GET("/",h.userIdenity,h.Home);
	v1 := router.Group("/user")
	{
		v1.POST("/signin",h.Signin)
		v1.POST("/signup",h.Signup)
		v1.POST("/reset-password",h.ResetPassword)
		v1.GET("/logout",h.Logout)
	}

	return router
}
