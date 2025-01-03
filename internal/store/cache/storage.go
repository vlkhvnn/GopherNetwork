package cache

import (
	"GopherNetwork/internal/store"
	"context"

	"github.com/go-redis/redis/v8"
)

type Storage struct {
	Users interface {
		Get(context.Context, int64) (*store.User, error)
		Set(context.Context, *store.User) error
	}
}

func NewRedisStorage(rbd *redis.Client) Storage {
	return Storage{
		Users: &UserStore{rdb: rbd},
	}
}
