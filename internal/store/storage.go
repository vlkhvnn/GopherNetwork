package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("record not found")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Post interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (*Post, error)
		DeleteById(context.Context, int64) error
		Update(context.Context, *Post) error
		GetUserFeed(context.Context, int64) ([]*PostMetaData, error)
	}
	User interface {
		Create(context.Context, *User) error
		GetById(context.Context, int64) (*User, error)
	}

	Comment interface {
		GetByPostId(context.Context, int64) ([]Comment, error)
		Create(context.Context, *Comment) error
	}

	Followers interface {
		Follow(context.Context, int64, int64) error
		Unfollow(context.Context, int64, int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Post:      &PostStore{db},
		User:      &UserStore{db},
		Comment:   &CommentStore{db},
		Followers: &FollowerStore{db},
	}
}
