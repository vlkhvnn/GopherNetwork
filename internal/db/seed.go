package db

import (
	"GopherNetwork/internal/store"
	"math/rand"
)

func generateFollowers(num int, users []*store.User) []*store.Follower {
	followers := make([]*store.Follower, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		follower := users[rand.Intn(len(users))]
		// Ensure a user does not follow themselves
		for user.ID == follower.ID {
			follower = users[rand.Intn(len(users))]
		}
		followers[i] = &store.Follower{
			UserId:     user.ID,
			FollowerId: follower.ID,
		}
	}
	return followers
}
