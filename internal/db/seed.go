package db

import (
	"GopherNetwork/internal/store"
	"context"
	"fmt"
	"log"
	"math/rand"
)

var usernames = []string{
	"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack",
	"Karen", "Leo", "Mona", "Nate", "Olivia", "Paul", "Quincy", "Rachel", "Steve", "Tina",
	"Uma", "Vince", "Wendy", "Xander", "Yara", "Zane", "Aria", "Blake", "Cora", "Dylan",
	"Elena", "Finn", "Gia", "Henry", "Isla", "Jake", "Kayla", "Liam", "Mia", "Noah",
	"Oscar", "Piper", "Quinn", "Ruby", "Sam", "Tara", "Ulysses", "Violet", "Willow", "Xena",
}

var titles = []string{
	"10 Tips for Better Sleep",
	"How to Start a Garden",
	"The Benefits of Daily Exercise",
	"Top 5 Programming Languages to Learn",
	"Easy Recipes for Busy Weeknights",
	"The Ultimate Travel Guide",
	"Understanding Cryptocurrency",
	"Mindfulness for Beginners",
	"How to Save Money on Groceries",
	"Best Practices for Remote Work",
	"Creating a Morning Routine",
	"How to Improve Your Photography Skills",
	"Top 10 Books to Read This Year",
	"Home Organization Hacks",
	"Learning a New Language",
	"DIY Home Improvement Projects",
	"The Importance of Hydration",
	"How to Build a Personal Brand",
	"Effective Time Management Strategies",
	"Tips for Healthy Eating",
}

var contents = []string{
	"Discover the secrets to getting a better night's sleep with these ten tips. From establishing a bedtime routine to creating a restful environment, learn how to improve your sleep quality.",
	"Starting a garden can be a rewarding and therapeutic experience. Learn the basics of gardening, from choosing the right plants to understanding soil and sunlight requirements.",
	"Daily exercise has numerous benefits for both physical and mental health. Explore the advantages of staying active and find easy ways to incorporate exercise into your daily routine.",
	"With so many programming languages available, it can be hard to know where to start. This guide highlights the top five programming languages to learn for a successful career in tech.",
	"Cooking doesn't have to be a chore. These easy recipes are perfect for busy weeknights, offering quick and delicious meals that the whole family will love.",
	"Planning your next adventure? This ultimate travel guide provides tips and recommendations for destinations around the world, helping you make the most of your trips.",
	"Cryptocurrency is a hot topic in the financial world. Learn the basics of cryptocurrency, how it works, and the potential benefits and risks of investing in digital currencies.",
	"Mindfulness can help reduce stress and improve overall well-being. Discover simple mindfulness practices that you can incorporate into your daily life to stay present and focused.",
	"Looking to cut down on your grocery bill? These money-saving tips will help you shop smarter and make the most of your food budget without sacrificing quality.",
	"Remote work is becoming increasingly common. Discover best practices for staying productive, maintaining work-life balance, and creating an effective home office setup.",
	"A well-planned morning routine can set the tone for a successful day. Learn how to create a morning routine that energizes you and helps you achieve your daily goals.",
	"Improve your photography skills with these tips and techniques. Whether you're a beginner or an experienced photographer, these insights will help you take better photos.",
	"Looking for your next great read? Check out this list of the top ten books to read this year, featuring a mix of fiction, non-fiction, and must-read classics.",
	"Keeping your home organized can be a challenge. These home organization hacks will help you declutter and create a more functional and tidy living space.",
	"Learning a new language can be a fun and rewarding experience. Explore different methods and resources to help you become fluent in a new language.",
	"DIY home improvement projects can save you money and add value to your home. Get inspired with these easy and affordable projects that you can tackle on your own.",
	"Staying hydrated is crucial for your health. Learn about the importance of hydration, how much water you should drink daily, and tips for increasing your water intake.",
	"Building a personal brand can open up new opportunities. Discover strategies for creating and promoting your personal brand, both online and offline.",
	"Time management is key to achieving your goals. These effective time management strategies will help you stay organized, prioritize tasks, and make the most of your time.",
	"Healthy eating doesn't have to be complicated. Get tips for maintaining a balanced diet, including meal planning ideas and nutritious recipes to keep you on track.",
}

var tags = []string{
	"Health", "Fitness", "Technology", "Travel", "Food", "Lifestyle", "Finance", "Education",
	"DIY", "Home Improvement", "Photography", "Books", "Mindfulness", "Gardening", "Recipes",
	"Remote Work", "Time Management", "Personal Development", "Hydration", "Language Learning",
}

var comments = []string{
	"Great post! Thanks for sharing.",
	"I found this very informative. Keep it up!",
	"Interesting read. I learned a lot.",
	"Thank you for the tips. Very useful!",
	"Excellent article. I appreciate the insights.",
	"This was very helpful. Thanks!",
	"Well written and informative. Thanks for posting.",
	"Thanks for the great advice!",
	"I really enjoyed this post. Keep them coming!",
	"Helpful and concise. Great job!",
	"Awesome post! I look forward to reading more.",
	"This was just what I needed. Thanks!",
	"Very good read. I will definitely try these tips.",
	"Thank you for the detailed information.",
	"I found this post very insightful.",
	"Great tips! I will definitely use them.",
	"Thanks for sharing your knowledge.",
	"I appreciate the practical advice in this post.",
	"This was an excellent read. Very informative.",
	"Thanks for the inspiration and ideas!",
}

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

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)
	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i+1),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i+1) + "@example.com",
			Role: store.Role{
				Name: "user",
			},
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			UserID:  user.ID,
			Content: contents[rand.Intn(len(contents))],
			Title:   titles[rand.Intn(len(titles))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}
	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		post := posts[rand.Intn(len(posts))]
		cms[i] = &store.Comment{
			PostID:  post.ID,
			UserID:  user.ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
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
