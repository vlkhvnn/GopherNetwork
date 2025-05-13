package cache

import (
	"GopherNetwork/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type UserStore struct {
	rdb *redis.Client
}

const UserExpTime = time.Hour * 24

func (s *UserStore) Get(ctx context.Context, userID int64) (*models.User, error) {
	cacheKey := fmt.Sprintf("user-%d", userID)
	data, err := s.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	var user models.User
	if data != "" {
		err := json.Unmarshal([]byte(data), &user)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}
func (s *UserStore) Set(ctx context.Context, user *models.User) error {
	cacheKey := fmt.Sprintf("user-%d", user.ID)
	json, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return s.rdb.SetEX(ctx, cacheKey, json, UserExpTime).Err()
}
