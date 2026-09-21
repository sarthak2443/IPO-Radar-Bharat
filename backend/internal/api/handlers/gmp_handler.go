package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/domain"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/service"
)

type GMPHandler struct {
	service service.GMPService
}

func NewGMPHandler(service service.GMPService) *GMPHandler {
	return &GMPHandler{service: service}
}

// GetLatestGMP handles GET /api/v1/ipos/:id/gmp
func (h *GMPHandler) GetLatestGMP(c echo.Context) error {
	ipoID := c.Param("id")
	gmp, err := h.service.GetLatestGMP(c.Request().Context(), ipoID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error":      "GMP data not found",
			"disclaimer": domain.UnofficialGMPDisclaimer,
		})
	}
	return c.JSON(http.StatusOK, gmp)
}

// GetGMPHistory handles GET /api/v1/ipos/:id/gmp/history
func (h *GMPHandler) GetGMPHistory(c echo.Context) error {
	ipoID := c.Param("id")
	history, err := h.service.GetGMPHistory(c.Request().Context(), ipoID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"ipoId":      ipoID,
		"disclaimer": domain.UnofficialGMPDisclaimer,
		"history":    history,
	})
}

// RecordGMP handles POST /api/v1/ipos/:id/gmp
func (h *GMPHandler) RecordGMP(c echo.Context) error {
	ipoID := c.Param("id")
	var gmp domain.GMPHistory
	if err := c.Bind(&gmp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	gmp.IPOID = ipoID
	if gmp.RecordedAt.IsZero() {
		gmp.RecordedAt = time.Now().UTC()
	}

	if err := h.service.RecordGMP(c.Request().Context(), &gmp); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, gmp)
}
