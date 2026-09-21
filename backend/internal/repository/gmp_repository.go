package repository

import (
	"context"

	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/domain"
	"gorm.io/gorm"
)

type GMPRepository interface {
	FindByIPOID(ctx context.Context, ipoID string) ([]domain.GMPHistory, error)
	GetLatest(ctx context.Context, ipoID string) (*domain.GMPHistory, error)
	Create(ctx context.Context, gmp *domain.GMPHistory) error
}

type gormGMPRepository struct {
	db *gorm.DB
}

func NewGMPRepository(db *gorm.DB) GMPRepository {
	return &gormGMPRepository{db: db}
}

func (r *gormGMPRepository) FindByIPOID(ctx context.Context, ipoID string) ([]domain.GMPHistory, error) {
	var list []domain.GMPHistory
	err := r.db.WithContext(ctx).
		Where("ipo_id = ?", ipoID).
		Order("recorded_at DESC").
		Find(&list).Error
	return list, err
}

func (r *gormGMPRepository) GetLatest(ctx context.Context, ipoID string) (*domain.GMPHistory, error) {
	var gmp domain.GMPHistory
	err := r.db.WithContext(ctx).
		Where("ipo_id = ?", ipoID).
		Order("recorded_at DESC").
		First(&gmp).Error
	if err != nil {
		return nil, err
	}
	return &gmp, nil
}

func (r *gormGMPRepository) Create(ctx context.Context, gmp *domain.GMPHistory) error {
	return r.db.WithContext(ctx).Create(gmp).Error
}
