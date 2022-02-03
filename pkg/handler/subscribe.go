package handler

import "github.com/gin-gonic/gin"

func (h *Handler) SubscribeTo(c *gin.Context) {
	id := c.Param("id"); 
	userid := GetUser(c);

	_,err := h.services.SubscribeTo(id,userid);
	



}	