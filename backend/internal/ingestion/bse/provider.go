package bse

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
	return "BSE"
}

// FetchIPOs fetches and normalizes IPOs from the BSE source.
func (p *Provider) FetchIPOs(ctx context.Context) ([]domain.IPO, error) {
	// TODO: Implement BSE IPO public issues scraping / API integration
	return []domain.IPO{}, nil
}
