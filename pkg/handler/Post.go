package handler

import (
	"Postresql/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Addpost(c *gin.Context){ 	
	var post structs.PostStruct;

	if err := c.BindJSON(&post);err != nil { 
		c.JSON(http.StatusBadRequest,gin.H {
			"errror" : err.Error(),
		})
		return
	}


}