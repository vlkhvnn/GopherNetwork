package store

import (
	"context"
	"database/sql"
	"time"
)

func NewMockStore() Storage {
	return Storage{
		User: &MockUserStorage{},
	}
}

type MockUserStorage struct{}

func (m *MockUserStorage) Create(ctx context.Context, t *sql.Tx, user *User) error {
	return nil
}

func (m *MockUserStorage) GetById(ctx context.Context, id int64) (*User, error) {
	return nil, nil
}

func (m *MockUserStorage) GetAll(ctx context.Context) ([]*User, error) {
	return nil, nil
}

func (m *MockUserStorage) GetByEmail(ctx context.Context, email string) (*User, error) {
	return nil, nil
}

func (m *MockUserStorage) CreateAndInvite(ctx context.Context, user *User, token string, timeDur time.Duration) error {
	return nil
}

func (m *MockUserStorage) Activate(ctx context.Context, token string) error {
	return nil
}

func (m *MockUserStorage) Delete(ctx context.Context, id int64) error {
	return nil
}
