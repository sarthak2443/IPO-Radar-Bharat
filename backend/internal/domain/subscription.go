package domain

import (
	"time"
)

type IPOSubscription struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	IPOID        string    `json:"ipoId" gorm:"type:varchar(64);index;not null"`
	Timestamp    time.Time `json:"timestamp" gorm:"index;not null"`
	Retail       float64   `json:"retail" gorm:"type:numeric(8,2);default:0"`
	QIB          float64   `json:"qib" gorm:"type:numeric(8,2);default:0"`
	NII          float64   `json:"nii" gorm:"type:numeric(8,2);default:0"`
	Employee     float64   `json:"employee" gorm:"type:numeric(8,2);default:0"`
	Shareholder  float64   `json:"shareholder" gorm:"type:numeric(8,2);default:0"`
	Total        float64   `json:"total" gorm:"type:numeric(8,2);default:0"`
	Applications int64     `json:"applications" gorm:"type:bigint;default:0"`
	CreatedAt    time.Time `json:"createdAt"`
}

func (IPOSubscription) TableName() string {
	return "ipo_subscriptions"
}
