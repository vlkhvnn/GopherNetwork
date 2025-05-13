package storage

import (
	"GopherNetwork/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type RoleStoreInterface interface {
	GetByName(ctx context.Context, name string) (*models.Role, error)
}

type RoleStore struct {
	db *gorm.DB
}

func NewRoleStore(db *gorm.DB) *RoleStore {
	return &RoleStore{db: db}
}

func (s *RoleStore) GetByName(ctx context.Context, name string) (*models.Role, error) {
	var role models.Role
	err := s.db.WithContext(ctx).Where("name = ?", name).First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &role, nil
}
