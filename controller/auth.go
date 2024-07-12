package controller

import (
	"net/http"
	"peterchu999/simple-api/utils/auth"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username, exist := c.Params.Get("username")
	if !exist {
		username = "default"
	}
	token, err := auth.GenerateToken(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
