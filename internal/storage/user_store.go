package storage

import (
	"GopherNetwork/internal/models"
	"context"
	"time"

	"gorm.io/gorm"
)

type UserStoreInterface interface {
	Create(context.Context, *models.User) (*models.User, error)
	GetByID(context.Context, int64) (*models.User, error)
	GetByEmail(context.Context, string) (*models.User, error)
	GetAll(context.Context) ([]*models.User, error)
	GetByUsername(context.Context, string) (*models.User, error)
	CreateAndInvite(context.Context, *models.User, string) error
	Activate(context.Context, string) error
	Delete(context.Context, int64) error
}

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) Create(ctx context.Context, user *models.User) (*models.User, error) {
	if err := s.db.WithContext(ctx).Create(user).Error; err != nil {
		// Handle unique constraint errors if needed
		return nil, err
	}
	return user, nil
}

func (s *UserStore) GetByID(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	if err := s.db.WithContext(ctx).Preload("Role").First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := s.db.WithContext(ctx).Preload("Role").Where("email = ? AND is_active = ?", email, true).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserStore) GetAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	if err := s.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserStore) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	if err := s.db.WithContext(ctx).Preload("Role").Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserStore) CreateAndInvite(ctx context.Context, user *models.User, token string) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		invitation := models.UserInvitation{
			Token:     token,
			UserID:    user.ID,
			ExpiresAt: time.Now().Add(24 * time.Hour),
		}

		if err := tx.Create(&invitation).Error; err != nil {
			return err
		}

		return nil
	})
}

func (s *UserStore) Activate(ctx context.Context, token string) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var invite models.UserInvitation
		if err := tx.Where("token = ? AND expires_at > ?", token, time.Now()).First(&invite).Error; err != nil {
			return ErrNotFound
		}

		if err := tx.Model(&models.User{}).Where("id = ?", invite.UserID).Update("is_active", true).Error; err != nil {
			return err
		}

		if err := tx.Delete(&models.UserInvitation{}, "user_id = ?", invite.UserID).Error; err != nil {
			return err
		}

		return nil
	})
}

func (s *UserStore) Delete(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.User{}, id).Error; err != nil {
			return err
		}
		if err := tx.Delete(&models.UserInvitation{}, "user_id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
}
