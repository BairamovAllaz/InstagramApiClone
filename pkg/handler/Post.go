package handler

import (
	"Postresql/structs"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Addpost(c *gin.Context) {
	var post structs.PostStruct

	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errror": err.Error(),
		})
		return
	}
	path, err := h.services.Addpost(post,c)
	if err != nil {
		fmt.Print("hey")
	}

	c.JSON(http.StatusOK, gin.H{
		"data": path,
	})

}
