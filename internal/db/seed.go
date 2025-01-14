package db

import (
	"GopherNetwork/internal/store"
	"context"
	"log"
	"math/rand"
)

func Seed(store store.Storage) {
	ctx := context.Background()

	// Assuming users are already created and fetched
	users, err := store.User.GetAll(ctx)
	if err != nil {
		log.Println("Error fetching users: ", err)
		return
	}

	// Generate and insert follower relationships
	followers := generateFollowers(300, users)
	for _, follower := range followers {
		if err := store.Followers.Follow(ctx, follower.UserId, follower.FollowerId); err != nil {
			log.Println("Error creating a follower: ", err)
		}
	}
	log.Println("Seeding Followers Completed")
}

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
