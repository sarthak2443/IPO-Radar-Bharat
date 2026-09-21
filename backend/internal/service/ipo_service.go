package service

import (
	"context"
	"errors"

	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/domain"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/repository"
)

var ErrDatabaseNotConnected = errors.New("database is not connected")

type IPOService interface {
	ListIPOs(ctx context.Context, status string) ([]domain.IPO, error)
	GetIPOByID(ctx context.Context, id string) (*domain.IPO, error)
	GetUpcomingIPOs(ctx context.Context) ([]domain.IPO, error)
	GetActiveIPOs(ctx context.Context) ([]domain.IPO, error)
	CreateIPO(ctx context.Context, ipo *domain.IPO) error
}

type ipoService struct {
	repo repository.IPORepository
}

func NewIPOService(repo repository.IPORepository) IPOService {
	return &ipoService{repo: repo}
}

func (s *ipoService) ListIPOs(ctx context.Context, status string) ([]domain.IPO, error) {
	if s.repo == nil {
		return []domain.IPO{}, nil
	}
	return s.repo.FindAll(ctx, status)
}

func (s *ipoService) GetIPOByID(ctx context.Context, id string) (*domain.IPO, error) {
	if s.repo == nil {
		return nil, ErrDatabaseNotConnected
	}
	return s.repo.FindByID(ctx, id)
}

func (s *ipoService) GetUpcomingIPOs(ctx context.Context) ([]domain.IPO, error) {
	if s.repo == nil {
		return []domain.IPO{}, nil
	}
	return s.repo.FindAll(ctx, string(domain.StatusUpcoming))
}

func (s *ipoService) GetActiveIPOs(ctx context.Context) ([]domain.IPO, error) {
	if s.repo == nil {
		return []domain.IPO{}, nil
	}
	return s.repo.FindAll(ctx, string(domain.StatusOpen))
}

func (s *ipoService) CreateIPO(ctx context.Context, ipo *domain.IPO) error {
	if s.repo == nil {
		return ErrDatabaseNotConnected
	}
	return s.repo.Create(ctx, ipo)
}
