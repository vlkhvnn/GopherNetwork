package storage

import (
	"GopherNetwork/internal/models"
	"context"

	"gorm.io/gorm"
)

type CommentStoreInterface interface {
	Create(context.Context, *models.Comment) (*models.Comment, error)
	GetByPostId(context.Context, int64) ([]models.Comment, error)
}

type CommentStore struct {
	db *gorm.DB
}

func NewCommentStore(db *gorm.DB) *CommentStore {
	return &CommentStore{db: db}
}

func (s *CommentStore) Create(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	if err := s.db.WithContext(ctx).Create(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *CommentStore) GetByPostId(ctx context.Context, postId int64) ([]models.Comment, error) {
	var comments []models.Comment
	err := s.db.WithContext(ctx).
		Preload("User").
		Where("post_id = ?", postId).
		Order("created_at desc").
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
