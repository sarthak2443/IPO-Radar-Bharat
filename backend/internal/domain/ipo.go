package domain

import (
	"time"
)

type IPOStatus string

const (
	StatusUpcoming IPOStatus = "upcoming"
	StatusOpen     IPOStatus = "open"
	StatusClosing  IPOStatus = "closing"
	StatusListed   IPOStatus = "listed"
	StatusClosed   IPOStatus = "closed"
)

type IPOCategory string

const (
	CategoryMainboard IPOCategory = "Mainboard"
	CategorySME       IPOCategory = "SME"
)

type IPO struct {
	ID                string      `json:"id" gorm:"primaryKey;type:varchar(64)"`
	CompanyName       string      `json:"companyName" gorm:"type:varchar(255);not null"`
	ShortName         string      `json:"shortName" gorm:"type:varchar(100)"`
	Symbol            string      `json:"symbol" gorm:"type:varchar(32);index"`
	ISIN              string      `json:"isin" gorm:"type:varchar(32);index"`
	Category          IPOCategory `json:"category" gorm:"type:varchar(32);default:'Mainboard'"`
	IPOType           string      `json:"ipoType" gorm:"type:varchar(32);default:'Book Built'"`
	Exchange          string      `json:"exchange" gorm:"type:varchar(32);default:'NSE/BSE'"`
	IssueSize         string      `json:"issueSize" gorm:"type:varchar(64)"`
	PriceBandLow      float64     `json:"priceBandLow" gorm:"type:numeric(12,2)"`
	PriceBandHigh     float64     `json:"priceBandHigh" gorm:"type:numeric(12,2)"`
	PriceBand         string      `json:"priceBand" gorm:"type:varchar(64)"`
	LotSize           int         `json:"lotSize" gorm:"type:integer"`
	MinimumInvestment string      `json:"minimumInvestment" gorm:"type:varchar(64)"`
	Sector            string      `json:"sector" gorm:"type:varchar(100)"`

	OpenDate      *time.Time `json:"openDate,omitempty" gorm:"index"`
	CloseDate     *time.Time `json:"closeDate,omitempty" gorm:"index"`
	AllotmentDate *time.Time `json:"allotmentDate,omitempty"`
	ListingDate   *time.Time `json:"listingDate,omitempty"`

	Status      IPOStatus `json:"status" gorm:"type:varchar(32);index;not null"`
	GMP         float64   `json:"gmp" gorm:"type:numeric(10,2);default:0"`
	Subscription float64  `json:"subscription" gorm:"type:numeric(10,2);default:0"`
	RadarScore  float64   `json:"radarScore" gorm:"type:numeric(5,2);default:0"`
	ListingGain float64   `json:"listingGain,omitempty" gorm:"type:numeric(6,2);default:0"`

	Subscriptions []IPOSubscription   `json:"subscriptions,omitempty" gorm:"foreignKey:IPOID"`
	GMPHistories  []GMPHistory        `json:"gmpHistories,omitempty" gorm:"foreignKey:IPOID"`
	Listing       *ListingPerformance `json:"listing,omitempty" gorm:"foreignKey:IPOID"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (IPO) TableName() string {
	return "ipos"
}
