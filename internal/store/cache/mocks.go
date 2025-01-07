package cache

import (
	"GopherNetwork/internal/store"
	"context"
)

func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStorage{},
	}
}

type MockUserStorage struct{}

func (s *MockUserStorage) Get(ctx context.Context, id int64) (*store.User, error) {
	return nil, nil
}

func (s *MockUserStorage) Set(ctx context.Context, user *store.User) error {
	return nil
}
