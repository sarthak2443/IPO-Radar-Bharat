package kafka

import "time"

// Event types defined in architecture roadmap
const (
	EventIPOCreated          = "ipo.created"
	EventIPOUpdated          = "ipo.updated"
	EventSubscriptionUpdated = "subscription.updated"
	EventIPOListed           = "ipo.listed"
	EventGMPUpdated          = "gmp.updated"
	EventFilingDiscovered    = "filing.discovered"
)

type EventEnvelope struct {
	Event     string      `json:"event"`
	IPOID     string      `json:"ipo_id"`
	Source    string      `json:"source"`
	Timestamp time.Time   `json:"timestamp"`
	Data      interface{} `json:"data"`
}

type SubscriptionUpdatedData struct {
	Retail       float64 `json:"retail"`
	QIB          float64 `json:"qib"`
	NII          float64 `json:"nii"`
	Employee     float64 `json:"employee,omitempty"`
	Shareholder  float64 `json:"shareholder,omitempty"`
	Total        float64 `json:"total"`
	Applications int64   `json:"applications,omitempty"`
}

type GMPUpdatedData struct {
	GMP        float64 `json:"gmp"`
	Source     string  `json:"source"`
	RecordedAt string  `json:"recorded_at"`
}
