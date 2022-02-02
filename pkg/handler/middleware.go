package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const Userid = "Userid";

func (h *Handler) userIdenity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized,gin.H{ 
			"error" : "invalid user header",
		})
		return
	}
	headerparts := strings.Split(header, " ")
	if len(headerparts) != 2 || headerparts[0] != "Bearer" {
		c.JSON(http.StatusBadRequest,gin.H { 
			"result" : "invalid auth header",
		})
		return
	}
	userId, err := h.services.Parsetoken(headerparts[1]);
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H { 
			"result" : "user not found",
		})
		return
	}
	c.Set(Userid, userId)
}



func GetUser(c *gin.Context) int{  
	id, _ := c.Get(Userid);
	data, _ := id.(int)
	return data;
}


