package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/domain"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/service"
)

type IPOHandler struct {
	service service.IPOService
}

func NewIPOHandler(service service.IPOService) *IPOHandler {
	return &IPOHandler{service: service}
}

// ListIPOs handles GET /api/v1/ipos?status=...
func (h *IPOHandler) ListIPOs(c echo.Context) error {
	status := c.QueryParam("status")
	ipos, err := h.service.ListIPOs(c.Request().Context(), status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, ipos)
}

// GetUpcomingIPOs handles GET /api/v1/ipos/upcoming
func (h *IPOHandler) GetUpcomingIPOs(c echo.Context) error {
	ipos, err := h.service.GetUpcomingIPOs(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, ipos)
}

// GetActiveIPOs handles GET /api/v1/ipos/active
func (h *IPOHandler) GetActiveIPOs(c echo.Context) error {
	ipos, err := h.service.GetActiveIPOs(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, ipos)
}

// GetIPOByID handles GET /api/v1/ipos/:id
func (h *IPOHandler) GetIPOByID(c echo.Context) error {
	id := c.Param("id")
	ipo, err := h.service.GetIPOByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "IPO not found"})
	}
	return c.JSON(http.StatusOK, ipo)
}

// CreateIPO handles POST /api/v1/ipos
func (h *IPOHandler) CreateIPO(c echo.Context) error {
	var ipo domain.IPO
	if err := c.Bind(&ipo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.service.CreateIPO(c.Request().Context(), &ipo); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, ipo)
}
