package handler

import (
	"Postresql/structs"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Addpost(c *gin.Context) {
	userid := GetUser(c);

	var post structs.PostStruct

	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errror": err.Error(),
		})
		return
	}
	path, err := h.services.Addpost(post,c,userid)
	if err != nil {
		fmt.Print("hey")
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success added post too see it follow link",
		"data": path,
	})
}

func(h *Handler)AddLikeToPost(c *gin.Context) { 

	id := c.Param("id"); 
	userid := GetUser(c);

	_,err := h.services.AddLikeToPostService(id,userid);	

	if err != nil { 
		c.JSON(http.StatusOK,gin.H{ 
			"result" : err.Error(),
		})
		return;
	}
	c.JSON(http.StatusOK,gin.H{ 
		"result" : "work",
	})
}

func(h *Handler)DeletPost(c *gin.Context) { 
	id := c.Param("id"); 
	userid := GetUser(c);
	_,err := h.services.DeletPost(id,userid);
	if err != nil { 
		c.JSON(http.StatusBadRequest,gin.H{ 
			"Error" : err.Error(),
		})
		return;
	}
	c.JSON(http.StatusOK,gin.H{  
		"Result" : "Delete post success",
	})
}


func(h *Handler)AddDisLikeToPost(c *gin.Context) { 
	id := c.Param("id"); 
	Userid := GetUser(c); 

	_,err := h.services.AddDislikeToPost(id,Userid);

	if err != nil { 
		c.JSON(http.StatusBadRequest,gin.H{  
			"error" : err.Error(),
		})
		return;
	}

	c.JSON(http.StatusOK,gin.H{  
		"result" : "you are ok!! dislike it!!!!",
	})
}
