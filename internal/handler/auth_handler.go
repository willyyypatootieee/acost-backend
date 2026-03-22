package handler

import (
	"acost/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CredentialsRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  string `json:"user"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	User    string `json:"user"`
}

type MeResponse struct {
	User string `json:"user"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// Login godoc
// @Summary Login user
// @Description Authenticates a user and returns a static bearer token for development.
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body CredentialsRequest true "Login payload"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} ErrorResponse
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var req CredentialsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": config.StaticAuthToken,
		"user":  req.Username,
	})
}

// Register godoc
// @Summary Register user
// @Description Registers a user and returns a static bearer token for development.
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body CredentialsRequest true "Register payload"
// @Success 201 {object} RegisterResponse
// @Failure 400 {object} ErrorResponse
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var req CredentialsRequest
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

// Me godoc
// @Summary Current user
// @Description Returns the current authenticated user from bearer token middleware.
// @Tags auth
// @Produce json
// @Security BearerAuth
// @Success 200 {object} MeResponse
// @Failure 401 {object} ErrorResponse
// @Router /me [get]
func Me(c *gin.Context) {
	user := c.GetString("user")
	if user == "" {
		user = "authenticated-user"
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
