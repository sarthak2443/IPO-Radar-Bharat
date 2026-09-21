package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/api/handlers"
)

type Handlers struct {
	Health       *handlers.HealthHandler
	IPO          *handlers.IPOHandler
	Subscription *handlers.SubscriptionHandler
	GMP          *handlers.GMPHandler
	Listing      *handlers.ListingHandler
}

func RegisterRoutes(e *echo.Echo, h *Handlers) {
	// Root health endpoints
	e.GET("/healthz", h.Health.HealthCheck)

	// API v1 group
	v1 := e.Group("/api/v1")
	{
		v1.GET("/health", h.Health.HealthCheck)

		// IPO routes
		ipos := v1.Group("/ipos")
		{
			ipos.GET("", h.IPO.ListIPOs)
			ipos.POST("", h.IPO.CreateIPO)
			ipos.GET("/upcoming", h.IPO.GetUpcomingIPOs)
			ipos.GET("/active", h.IPO.GetActiveIPOs)
			ipos.GET("/:id", h.IPO.GetIPOByID)

			// Subscription routes
			ipos.GET("/:id/subscription", h.Subscription.GetLatestSubscription)
			ipos.POST("/:id/subscription", h.Subscription.RecordSubscription)
			ipos.GET("/:id/subscription/history", h.Subscription.GetSubscriptionHistory)

			// GMP routes (unofficial)
			ipos.GET("/:id/gmp", h.GMP.GetLatestGMP)
			ipos.POST("/:id/gmp", h.GMP.RecordGMP)
			ipos.GET("/:id/gmp/history", h.GMP.GetGMPHistory)

			// Listing performance
			ipos.GET("/:id/listing", h.Listing.GetListingPerformance)
		}
	}
}
