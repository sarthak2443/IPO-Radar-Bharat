package service

import (
	"context"

	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/domain"
	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/repository"
)

type SubscriptionService interface {
	GetSubscriptionHistory(ctx context.Context, ipoID string) ([]domain.IPOSubscription, error)
	GetLatestSubscription(ctx context.Context, ipoID string) (*domain.IPOSubscription, error)
	RecordSubscription(ctx context.Context, sub *domain.IPOSubscription) error
}

type subscriptionService struct {
	repo repository.SubscriptionRepository
}

func NewSubscriptionService(repo repository.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{repo: repo}
}

func (s *subscriptionService) GetSubscriptionHistory(ctx context.Context, ipoID string) ([]domain.IPOSubscription, error) {
	return s.repo.FindByIPOID(ctx, ipoID)
}

func (s *subscriptionService) GetLatestSubscription(ctx context.Context, ipoID string) (*domain.IPOSubscription, error) {
	return s.repo.GetLatest(ctx, ipoID)
}

func (s *subscriptionService) RecordSubscription(ctx context.Context, sub *domain.IPOSubscription) error {
	return s.repo.Create(ctx, sub)
}
