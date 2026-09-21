package repository

import (
	"context"

	"github.com/sarthak2443/ipo-radar-bharat/backend/internal/domain"
	"gorm.io/gorm"
)

type IPORepository interface {
	FindAll(ctx context.Context, status string) ([]domain.IPO, error)
	FindByID(ctx context.Context, id string) (*domain.IPO, error)
	Create(ctx context.Context, ipo *domain.IPO) error
	Update(ctx context.Context, ipo *domain.IPO) error
	Delete(ctx context.Context, id string) error
}

type gormIPORepository struct {
	db *gorm.DB
}

func NewIPORepository(db *gorm.DB) IPORepository {
	return &gormIPORepository{db: db}
}

func (r *gormIPORepository) FindAll(ctx context.Context, status string) ([]domain.IPO, error) {
	var ipos []domain.IPO
	query := r.db.WithContext(ctx).Order("open_date DESC")
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Find(&ipos).Error
	return ipos, err
}

func (r *gormIPORepository) FindByID(ctx context.Context, id string) (*domain.IPO, error) {
	var ipo domain.IPO
	err := r.db.WithContext(ctx).
		Preload("Subscriptions").
		Preload("GMPHistories").
		Preload("Listing").
		Where("id = ?", id).
		First(&ipo).Error
	if err != nil {
		return nil, err
	}
	return &ipo, nil
}

func (r *gormIPORepository) Create(ctx context.Context, ipo *domain.IPO) error {
	return r.db.WithContext(ctx).Create(ipo).Error
}

func (r *gormIPORepository) Update(ctx context.Context, ipo *domain.IPO) error {
	return r.db.WithContext(ctx).Save(ipo).Error
}

func (r *gormIPORepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&domain.IPO{}, "id = ?", id).Error
}
