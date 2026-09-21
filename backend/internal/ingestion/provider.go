package ingestion

import (
	"context"

	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/domain"
)

// IPOProvider defines the contract for fetching IPO data from exchanges and regulatory sources.
type IPOProvider interface {
	Name() string
	FetchIPOs(ctx context.Context) ([]domain.IPO, error)
}

// SubscriptionProvider defines the contract for fetching live subscription data.
type SubscriptionProvider interface {
	FetchSubscription(ctx context.Context, ipoID string) (*domain.IPOSubscription, error)
}
