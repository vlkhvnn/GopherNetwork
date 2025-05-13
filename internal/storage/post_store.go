package storage

import (
	"GopherNetwork/internal/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type PostStoreInterface interface {
	Create(context.Context, *models.Post) error
	GetById(context.Context, int64) (*models.Post, error)
	DeleteById(context.Context, int64) error
	Update(context.Context, *models.Post) error
	GetUserFeed(context.Context, int64, PaginatedFeedQuery) ([]models.Post, error)
}

type PostStore struct {
	db *gorm.DB
}

func NewPostStore(db *gorm.DB) *PostStore {
	return &PostStore{db: db}
}

func (s *PostStore) Create(ctx context.Context, post *models.Post) error {
	if err := s.db.WithContext(ctx).Create(post).Error; err != nil {
		return err
	}
	return nil
}

func (s *PostStore) GetById(ctx context.Context, id int64) (*models.Post, error) {
	var post models.Post
	if err := s.db.WithContext(ctx).
		Preload("User").
		Preload("Comments").
		First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &post, nil
}

func (s *PostStore) DeleteById(ctx context.Context, id int64) error {
	result := s.db.WithContext(ctx).Delete(&models.Post{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (s *PostStore) Update(ctx context.Context, post *models.Post) error {
	// Optimistic locking not directly available, but simulate with a version check if needed
	result := s.db.WithContext(ctx).
		Model(&models.Post{}).
		Where("id = ?", post.ID).
		Updates(map[string]interface{}{
			"title":      post.Title,
			"content":    post.Content,
			"updated_at": post.UpdatedAt,
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (s *PostStore) GetUserFeed(ctx context.Context, userID int64, fq PaginatedFeedQuery) ([]models.Post, error) {
	var posts []models.Post

	tx := s.db.WithContext(ctx).
		Joins("JOIN followers f ON f.user_id = posts.user_id").
		Where("f.follower_id = ?", userID)

	if fq.Search != "" {
		tx = tx.Where("posts.title ILIKE ? OR posts.content ILIKE ?", "%"+fq.Search+"%", "%"+fq.Search+"%")
	}

	if len(fq.Tags) > 0 {
		// Assuming tags are stored as a string array in the DB (PostgreSQL)
		// You'll need to implement this logic depending on your tags model
		// You could consider using GORM's `clause.Expr` or a separate tag table
		// Placeholder:
		fmt.Println("Tag filtering requires implementation based on schema.")
	}

	if fq.Sort != "asc" && fq.Sort != "desc" {
		fq.Sort = "desc"
	}

	tx = tx.Preload("User").
		Order("posts.created_at " + fq.Sort).
		Limit(fq.Limit).
		Offset(fq.Offset)

	if err := tx.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
