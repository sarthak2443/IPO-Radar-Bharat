package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/config"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/ingestion/nse"
)

func main() {
	log.Println("⚙️  Starting IPO Radar Bharat Background Ingestion Worker...")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load worker configuration: %v", err)
	}

	nseProvider := nse.NewProvider()
	log.Printf("Loaded provider: %s (Environment: %s)\n", nseProvider.Name(), cfg.AppEnv)

	// Context with cancellation for graceful worker shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Ingestion loop ticker (e.g. periodic check)
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case t := <-ticker.C:
				log.Printf("Ingestion worker heartbeat at %s\n", t.UTC().Format(time.RFC3339))
				// Future: trigger scheduled ingestion tasks
			}
		}
	}()

	<-quit
	log.Println("🛑 Worker stopping...")
	cancel()
	log.Println("👋 Worker stopped successfully")
}
