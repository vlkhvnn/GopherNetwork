package storage

import (
	"GopherNetwork/internal/models"
	"context"

	"gorm.io/gorm"
)

type FollowerStoreInterface interface {
	Follow(context.Context, int64, int64) error
	Unfollow(context.Context, int64, int64) error
}

type FollowerStore struct {
	db *gorm.DB
}

func NewFollowerStore(db *gorm.DB) *FollowerStore {
	return &FollowerStore{db: db}
}

func (s *FollowerStore) Follow(ctx context.Context, followerID, followingID int64) error {
	f := &models.Follower{
		FollowerID:  followerID,
		FollowingID: followingID,
	}
	return s.db.WithContext(ctx).Create(f).Error
}

func (s *FollowerStore) Unfollow(ctx context.Context, followerID, followingID int64) error {
	return s.db.WithContext(ctx).
		Where("follower_id = ? AND following_id = ?", followerID, followingID).
		Delete(&models.Follower{}).Error
}
