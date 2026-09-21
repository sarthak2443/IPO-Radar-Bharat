package domain

import (
	"time"
)

const UnofficialGMPDisclaimer = "Grey Market Premium (GMP) is an unofficial, unregulated market metric. It is not provided or endorsed by NSE, BSE, or SEBI. For information purposes only."

type GMPHistory struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	IPOID      string    `json:"ipoId" gorm:"type:varchar(64);index;not null"`
	GMP        float64   `json:"gmp" gorm:"type:numeric(10,2);not null"`
	Source     string    `json:"source" gorm:"type:varchar(100);not null"`
	Disclaimer string    `json:"disclaimer" gorm:"type:varchar(255);default:'Unofficial market indicator'"`
	RecordedAt time.Time `json:"recordedAt" gorm:"index;not null"`
	CreatedAt  time.Time `json:"createdAt"`
}

func (GMPHistory) TableName() string {
	return "gmp_histories"
}
