package repository

import (
	"context"

	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/domain"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	FindByIPOID(ctx context.Context, ipoID string) ([]domain.IPOSubscription, error)
	GetLatest(ctx context.Context, ipoID string) (*domain.IPOSubscription, error)
	Create(ctx context.Context, sub *domain.IPOSubscription) error
}

type gormSubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &gormSubscriptionRepository{db: db}
}

func (r *gormSubscriptionRepository) FindByIPOID(ctx context.Context, ipoID string) ([]domain.IPOSubscription, error) {
	var subs []domain.IPOSubscription
	err := r.db.WithContext(ctx).
		Where("ipo_id = ?", ipoID).
		Order("timestamp DESC").
		Find(&subs).Error
	return subs, err
}

func (r *gormSubscriptionRepository) GetLatest(ctx context.Context, ipoID string) (*domain.IPOSubscription, error) {
	var sub domain.IPOSubscription
	err := r.db.WithContext(ctx).
		Where("ipo_id = ?", ipoID).
		Order("timestamp DESC").
		First(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *gormSubscriptionRepository) Create(ctx context.Context, sub *domain.IPOSubscription) error {
	return r.db.WithContext(ctx).Create(sub).Error
}
