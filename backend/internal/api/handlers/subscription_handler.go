package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/domain"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/service"
)

type SubscriptionHandler struct {
	service service.SubscriptionService
}

func NewSubscriptionHandler(service service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{service: service}
}

// GetLatestSubscription handles GET /api/v1/ipos/:id/subscription
func (h *SubscriptionHandler) GetLatestSubscription(c echo.Context) error {
	ipoID := c.Param("id")
	sub, err := h.service.GetLatestSubscription(c.Request().Context(), ipoID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Subscription data not found"})
	}
	return c.JSON(http.StatusOK, sub)
}

// GetSubscriptionHistory handles GET /api/v1/ipos/:id/subscription/history
func (h *SubscriptionHandler) GetSubscriptionHistory(c echo.Context) error {
	ipoID := c.Param("id")
	subs, err := h.service.GetSubscriptionHistory(c.Request().Context(), ipoID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, subs)
}

// RecordSubscription handles POST /api/v1/ipos/:id/subscription
func (h *SubscriptionHandler) RecordSubscription(c echo.Context) error {
	ipoID := c.Param("id")
	var sub domain.IPOSubscription
	if err := c.Bind(&sub); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	sub.IPOID = ipoID
	if sub.Timestamp.IsZero() {
		sub.Timestamp = time.Now().UTC()
	}

	if err := h.service.RecordSubscription(c.Request().Context(), &sub); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, sub)
}
