package nse

import (
	"context"
	"net/http"
	"time"

	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/domain"
)

type Provider struct {
	client *http.Client
}

func NewProvider() *Provider {
	return &Provider{
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (p *Provider) Name() string {
	return "NSE"
}

// FetchIPOs fetches and normalizes IPOs from the NSE source.
// Milestone 3 implementation placeholder.
func (p *Provider) FetchIPOs(ctx context.Context) ([]domain.IPO, error) {
	// TODO: Implement actual NSE API fetching, session cookie handling, and response parsing
	return []domain.IPO{}, nil
}

// FetchSubscription fetches live subscription figures for a given IPO.
// Milestone 3 implementation placeholder.
func (p *Provider) FetchSubscription(ctx context.Context, ipoID string) (*domain.IPOSubscription, error) {
	// TODO: Implement actual NSE live bid subscription fetching
	return nil, nil
}
