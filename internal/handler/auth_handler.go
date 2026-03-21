package handler

import (
	"acost/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req credentials
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": config.StaticAuthToken,
		"user":  req.Username,
	})
}

func Register(c *gin.Context) {
	var req credentials
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered",
		"token":   config.StaticAuthToken,
		"user":    req.Username,
	})
}

func Me(c *gin.Context) {
	user := c.GetString("user")
	if user == "" {
		user = "authenticated-user"
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
