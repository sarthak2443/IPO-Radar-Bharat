package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/config"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	logLevel := logger.Info
	if cfg.AppEnv == "production" {
		logLevel = logger.Warn
	}

	db, err := gorm.Open(postgres.Open(cfg.DBURL), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB handle: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	log.Println("Running database auto-migrations...")
	return db.AutoMigrate(
		&domain.IPO{},
		&domain.IPOSubscription{},
		&domain.GMPHistory{},
		&domain.ListingPerformance{},
	)
}
