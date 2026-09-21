package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "healthy",
		"service":   "ipo-radar-bharat-api",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}
