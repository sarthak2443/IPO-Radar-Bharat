package sebi

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
	return "SEBI"
}

// FetchIPOs fetches regulatory filings (DRHP, RHP, Prospectus) from SEBI public registry.
func (p *Provider) FetchIPOs(ctx context.Context) ([]domain.IPO, error) {
	// TODO: Implement SEBI public filings discovery
	return []domain.IPO{}, nil
}
