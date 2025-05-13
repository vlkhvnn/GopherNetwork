package storage

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrNotFound          = errors.New("record not found")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	db            *gorm.DB
	UserStore     UserStoreInterface
	CommentStore  CommentStoreInterface
	FollowerStore FollowerStoreInterface
	PostStore     PostStoreInterface
	RoleStore     RoleStoreInterface
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{
		db:            db,
		UserStore:     NewUserStore(db),
		CommentStore:  NewCommentStore(db),
		FollowerStore: NewFollowerStore(db),
		PostStore:     NewPostStore(db),
		RoleStore:     NewRoleStore(db),
	}
}
