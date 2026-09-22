package main

import (
	"context"
	"log"
	"time"

	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/config"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/domain"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/repository"
	"gorm.io/gorm/clause"
)

func parseDate(d string) *time.Time {
	t, err := time.Parse("2006-01-02", d)
	if err != nil {
		return nil
	}
	return &t
}

func main() {
	log.Println("Starting database seeding with real Indian IPO records...")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := repository.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v. Make sure postgres container is running via 'docker compose up -d'.", err)
	}

	if err := repository.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Real Indian IPO records
	ipos := []domain.IPO{
		{
			ID:                "hyundai-motor-india",
			CompanyName:       "Hyundai Motor India Ltd",
			ShortName:         "Hyundai Motor",
			Symbol:            "HYUNDAI",
			ISIN:              "INE000000001",
			Category:          domain.CategoryMainboard,
			IPOType:           "Book Built",
			Exchange:          "NSE/BSE",
			IssueSize:         "₹27,870 Cr",
			PriceBandLow:      1865,
			PriceBandHigh:     1960,
			PriceBand:         "₹1,865 – ₹1,960",
			LotSize:           7,
			MinimumInvestment: "₹13,720",
			Sector:            "Automobile",
			OpenDate:          parseDate("2024-10-15"),
			CloseDate:         parseDate("2024-10-17"),
			ListingDate:       parseDate("2024-10-22"),
			Status:            domain.StatusListed,
			GMP:               -25,
			Subscription:      2.37,
			RadarScore:        84,
			ListingGain:       -1.32,
		},
		{
			ID:                "swiggy-ltd",
			CompanyName:       "Swiggy Ltd",
			ShortName:         "Swiggy",
			Symbol:            "SWIGGY",
			ISIN:              "INE000000002",
			Category:          domain.CategoryMainboard,
			IPOType:           "Book Built",
			Exchange:          "NSE/BSE",
			IssueSize:         "₹11,327 Cr",
			PriceBandLow:      371,
			PriceBandHigh:     390,
			PriceBand:         "₹371 – ₹390",
			LotSize:           38,
			MinimumInvestment: "₹14,820",
			Sector:            "Consumer Tech",
			OpenDate:          parseDate("2024-11-06"),
			CloseDate:         parseDate("2024-11-08"),
			ListingDate:       parseDate("2024-11-13"),
			Status:            domain.StatusListed,
			GMP:               18,
			Subscription:      3.59,
			RadarScore:        89,
			ListingGain:       7.69,
		},
		{
			ID:                "ntpc-green-energy",
			CompanyName:       "NTPC Green Energy Ltd",
			ShortName:         "NTPC Green",
			Symbol:            "NTPCGREEN",
			ISIN:              "INE000000003",
			Category:          domain.CategoryMainboard,
			IPOType:           "Book Built",
			Exchange:          "NSE/BSE",
			IssueSize:         "₹10,000 Cr",
			PriceBandLow:      102,
			PriceBandHigh:     108,
			PriceBand:         "₹102 – ₹108",
			LotSize:           138,
			MinimumInvestment: "₹14,904",
			Sector:            "Clean Energy",
			OpenDate:          parseDate("2024-11-19"),
			CloseDate:         parseDate("2024-11-22"),
			ListingDate:       parseDate("2024-11-27"),
			Status:            domain.StatusListed,
			GMP:               3.5,
			Subscription:      2.42,
			RadarScore:        82,
			ListingGain:       3.24,
		},
		{
			ID:                "waaree-energies",
			CompanyName:       "Waaree Energies Ltd",
			ShortName:         "Waaree Energies",
			Symbol:            "WAAREE",
			ISIN:              "INE000000004",
			Category:          domain.CategoryMainboard,
			IPOType:           "Book Built",
			Exchange:          "NSE/BSE",
			IssueSize:         "₹4,321 Cr",
			PriceBandLow:      1427,
			PriceBandHigh:     1503,
			PriceBand:         "₹1,427 – ₹1,503",
			LotSize:           9,
			MinimumInvestment: "₹13,527",
			Sector:            "Renewable Energy",
			OpenDate:          parseDate("2024-10-21"),
			CloseDate:         parseDate("2024-10-23"),
			ListingDate:       parseDate("2024-10-28"),
			Status:            domain.StatusListed,
			GMP:               1550,
			Subscription:      76.34,
			RadarScore:        96,
			ListingGain:       69.66,
		},
		{
			ID:                "bajaj-housing-finance",
			CompanyName:       "Bajaj Housing Finance Ltd",
			ShortName:         "Bajaj Housing",
			Symbol:            "BAJAJHFL",
			ISIN:              "INE000000005",
			Category:          domain.CategoryMainboard,
			IPOType:           "Book Built",
			Exchange:          "NSE/BSE",
			IssueSize:         "₹6,560 Cr",
			PriceBandLow:      66,
			PriceBandHigh:     70,
			PriceBand:         "₹66 – ₹70",
			LotSize:           214,
			MinimumInvestment: "₹14,980",
			Sector:            "Financial Services",
			OpenDate:          parseDate("2024-09-09"),
			CloseDate:         parseDate("2024-09-11"),
			ListingDate:       parseDate("2024-09-16"),
			Status:            domain.StatusListed,
			GMP:               82,
			Subscription:      63.61,
			RadarScore:        98,
			ListingGain:       114.28,
		},
		{
			ID:                "tata-capital-ltd",
			CompanyName:       "Tata Capital Ltd",
			ShortName:         "Tata Capital",
			Symbol:            "TATACAP",
			ISIN:              "INE000000006",
			Category:          domain.CategoryMainboard,
			IPOType:           "Book Built",
			Exchange:          "NSE/BSE",
			IssueSize:         "₹15,000 Cr",
			PriceBandLow:      0,
			PriceBandHigh:     0,
			PriceBand:         "To be announced",
			LotSize:           0,
			MinimumInvestment: "To be announced",
			Sector:            "Financial Services",
			OpenDate:          parseDate("2026-10-15"),
			CloseDate:         parseDate("2026-10-18"),
			ListingDate:       parseDate("2026-10-25"),
			Status:            domain.StatusUpcoming,
			GMP:               120,
			Subscription:      0,
			RadarScore:        94,
		},
		{
			ID:                "ather-energy-ltd",
			CompanyName:       "Ather Energy Ltd",
			ShortName:         "Ather Energy",
			Symbol:            "ATHER",
			ISIN:              "INE000000007",
			Category:          domain.CategoryMainboard,
			IPOType:           "Book Built",
			Exchange:          "NSE/BSE",
			IssueSize:         "₹3,100 Cr",
			PriceBandLow:      0,
			PriceBandHigh:     0,
			PriceBand:         "To be announced",
			LotSize:           0,
			MinimumInvestment: "To be announced",
			Sector:            "Electric Vehicles",
			OpenDate:          parseDate("2026-10-28"),
			CloseDate:         parseDate("2026-10-31"),
			ListingDate:       parseDate("2026-11-06"),
			Status:            domain.StatusUpcoming,
			GMP:               65,
			Subscription:      0,
			RadarScore:        87,
		},
		{
			ID:                "bharat-innovations-sme",
			CompanyName:       "Bharat Precision Tech Ltd",
			ShortName:         "Bharat Precision",
			Symbol:            "BHARATPRE",
			ISIN:              "INE000000008",
			Category:          domain.CategorySME,
			IPOType:           "Fixed Price",
			Exchange:          "NSE SME",
			IssueSize:         "₹48 Cr",
			PriceBandLow:      115,
			PriceBandHigh:     122,
			PriceBand:         "₹115 – ₹122",
			LotSize:           1000,
			MinimumInvestment: "₹1,22,000",
			Sector:            "Engineering",
			OpenDate:          parseDate("2026-09-20"),
			CloseDate:         parseDate("2026-09-24"),
			ListingDate:       parseDate("2026-09-29"),
			Status:            domain.StatusOpen,
			GMP:               45,
			Subscription:      8.45,
			RadarScore:        78,
		},
	}

	for _, ipo := range ipos {
		err := db.WithContext(context.Background()).Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&ipo).Error
		if err != nil {
			log.Printf("Failed to seed IPO %s: %v", ipo.CompanyName, err)
		} else {
			log.Printf("Seeded IPO: %s (%s)", ipo.CompanyName, ipo.Status)
		}
	}

	log.Printf("Successfully seeded %d IPO records into PostgreSQL database.", len(ipos))
}
