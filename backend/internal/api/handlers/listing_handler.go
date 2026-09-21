package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/service"
)

type ListingHandler struct {
	ipoService service.IPOService
}

func NewListingHandler(ipoService service.IPOService) *ListingHandler {
	return &ListingHandler{ipoService: ipoService}
}

// GetListingPerformance handles GET /api/v1/ipos/:id/listing
func (h *ListingHandler) GetListingPerformance(c echo.Context) error {
	ipoID := c.Param("id")
	ipo, err := h.ipoService.GetIPOByID(c.Request().Context(), ipoID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "IPO not found"})
	}

	if ipo.Listing == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Listing performance data not yet available for this IPO"})
	}

	return c.JSON(http.StatusOK, ipo.Listing)
}
