package domain

import (
	"time"
)

type ListingPerformance struct {
	ID                 uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	IPOID              string    `json:"ipoId" gorm:"type:varchar(64);uniqueIndex;not null"`
	IssuePrice         float64   `json:"issuePrice" gorm:"type:numeric(12,2);not null"`
	ListingPrice       float64   `json:"listingPrice" gorm:"type:numeric(12,2);not null"`
	ListingGainPercent float64   `json:"listingGainPercent" gorm:"type:numeric(8,2);not null"`
	CurrentPrice       float64   `json:"currentPrice" gorm:"type:numeric(12,2)"`
	CurrentGainPercent float64   `json:"currentGainPercent" gorm:"type:numeric(8,2)"`
	RecordedAt         time.Time `json:"recordedAt" gorm:"not null"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}

func (ListingPerformance) TableName() string {
	return "listing_performances"
}
