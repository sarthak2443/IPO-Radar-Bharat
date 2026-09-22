package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/api/handlers"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/api/middleware"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/api/routes"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/config"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/repository"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/service"
	"gorm.io/gorm"
)

func main() {
	log.Println("Initializing IPO Radar Bharat API server...")

	// 1. Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// 2. Connect to Database (GORM / Postgres)
	var db *gorm.DB
	db, err = repository.NewDB(cfg)
	if err != nil {
		log.Printf("Warning: Database connection failed: %v. Running in offline/fallback mode.", err)
	} else {
		log.Println("Connected to PostgreSQL database successfully.")
		// Run auto migrations
		if err := repository.AutoMigrate(db); err != nil {
			log.Printf("Warning: AutoMigrate error: %v", err)
		} else {
			log.Println("Database migrations applied successfully.")
		}
	}

	// 3. Initialize Repositories and Services
	var ipoRepo repository.IPORepository
	var subRepo repository.SubscriptionRepository
	var gmpRepo repository.GMPRepository

	if db != nil {
		ipoRepo = repository.NewIPORepository(db)
		subRepo = repository.NewSubscriptionRepository(db)
		gmpRepo = repository.NewGMPRepository(db)
	}

	ipoService := service.NewIPOService(ipoRepo)
	subService := service.NewSubscriptionService(subRepo)
	gmpService := service.NewGMPService(gmpRepo)

	// 4. Initialize Handlers
	h := &routes.Handlers{
		Health:       handlers.NewHealthHandler(),
		IPO:          handlers.NewIPOHandler(ipoService),
		Subscription: handlers.NewSubscriptionHandler(subService),
		GMP:          handlers.NewGMPHandler(gmpService),
		Listing:      handlers.NewListingHandler(ipoService),
	}

	// 5. Initialize Echo router & middleware
	e := echo.New()
	e.HideBanner = true
	middleware.SetupMiddleware(e)

	// 6. Register routes
	routes.RegisterRoutes(e, h)

	// 7. Start server with Graceful Shutdown
	go func() {
		addr := fmt.Sprintf(":%s", cfg.Port)
		log.Printf("IPO Radar Bharat API listening on http://localhost:%s\n", cfg.Port)
		if err := e.Start(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting HTTP server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
