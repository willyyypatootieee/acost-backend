package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

type VersionResponse struct {
	Version string `json:"version"`
}

// HealthCheck godoc
// @Summary Health check
// @Description Returns service health status.
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func HealthCheck(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "acost-backend",
	})
}

// VersionCheck godoc
// @Summary Service version
// @Description Returns current service version.
// @Tags health
// @Produce json
// @Success 200 {object} VersionResponse
// @Router /version [get]
func VersionCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "1.0.0",
	})
}
