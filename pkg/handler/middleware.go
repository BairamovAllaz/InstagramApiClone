package handler

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) userIdenity(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		log.Fatal("Invalid user header");
		return
	}

	headerparts := strings.Split(header, " ")
	if len(headerparts) != 2 || headerparts[0] != "Bearer" {
		log.Fatalf("invalid auth header")
		return
	}

	UserId, err := h.services.Parsetoken(headerparts[1]);
	if err != nil {
		log.Fatalf("user not found")
		return
	}
	c.Set("Userid", UserId)
}




