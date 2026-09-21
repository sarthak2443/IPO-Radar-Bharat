package service

import (
	"context"

	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/domain"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/repository"
)

type GMPService interface {
	GetGMPHistory(ctx context.Context, ipoID string) ([]domain.GMPHistory, error)
	GetLatestGMP(ctx context.Context, ipoID string) (*domain.GMPHistory, error)
	RecordGMP(ctx context.Context, gmp *domain.GMPHistory) error
}

type gmpService struct {
	repo repository.GMPRepository
}

func NewGMPService(repo repository.GMPRepository) GMPService {
	return &gmpService{repo: repo}
}

func (s *gmpService) GetGMPHistory(ctx context.Context, ipoID string) ([]domain.GMPHistory, error) {
	return s.repo.FindByIPOID(ctx, ipoID)
}

func (s *gmpService) GetLatestGMP(ctx context.Context, ipoID string) (*domain.GMPHistory, error) {
	return s.repo.GetLatest(ctx, ipoID)
}

func (s *gmpService) RecordGMP(ctx context.Context, gmp *domain.GMPHistory) error {
	if gmp.Disclaimer == "" {
		gmp.Disclaimer = domain.UnofficialGMPDisclaimer
	}
	return s.repo.Create(ctx, gmp)
}
