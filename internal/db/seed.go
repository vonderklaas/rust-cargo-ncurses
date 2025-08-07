package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand/v2"
	"social/internal/store"
)

var usernames = []string{
	"nick", "tim", "xander", "alex", "brooke", "cameron", "dylan", "emily", "finley", "grace",
	"hudson", "isabel", "jack", "kylie", "liam", "mia", "nathan", "olivia", "patrick", "quinn",
	"ryan", "sophia", "tyler", "ursula", "victor", "wendy", "xavier", "yara", "zane",
}

var titles = []string{
	"Quick Tips for Better Sleep", "Mastering Your Morning Routine",
	"The Art of Simple Living", "Boost Your Productivity Today",
	"Healthy Habits Made Easy", "Unlocking Your Creative Flow",
	"Budget Travel Hacks", "Mindfulness in Minutes",
	"Declutter Your Digital Life", "Photography Basics for Beginners",
	"Cooking with Five Ingredients", "Sustainable Living Starts Here",
	"Your Guide to Home Workouts", "Gardening for Small Spaces",
	"Reading More, Stressing Less", "Finding Your Next Hobby",
	"Build a Better Reading List", "The Power of Positive Thinking",
	"Simple Steps to Save Money", "Tech Gadgets You Need Now",
}

var contents = []string{
	"Discover actionable tips to improve your sleep quality starting tonight. From optimizing your bedroom environment to evening routines, we'll guide you to more restful nights.",
	"Kickstart your day with an effective morning routine. Learn how to structure your first few hours for maximum productivity and a positive mindset.",
	"Embrace the beauty of simple living. We explore how decluttering your life and focusing on essentials can bring peace and clarity.",
	"Unlock your full potential with these productivity hacks. We'll share strategies to manage your time better and achieve more with less effort.",
	"Cultivate wellness effortlessly. This guide breaks down easy-to-adopt healthy habits that can transform your physical and mental well-being.",
	"Break through creative blocks and find your flow. Explore techniques to spark inspiration and unleash your inner artist.",
	"Travel more without breaking the bank. Our top tips for finding cheap flights, accommodation, and experiences will make your dream trips a reality.",
	"Integrate mindfulness into your daily life, even when you're busy. Discover simple exercises to reduce stress and increase focus.",
	"Take control of your digital life. Learn how to organize your files, manage notifications, and reduce screen time for a more balanced existence.",
	"Step into the world of photography with confidence. This beginner's guide covers essential camera settings, composition, and editing tips.",
	"Whip up delicious meals with minimal ingredients. These recipes prove that you don't need a pantry full of items to create culinary masterpieces.",
	"Make eco-conscious choices every day. We'll show you practical ways to reduce your environmental footprint and live more sustainably.",
	"Transform your home into your personal gym. Discover effective workouts and essential equipment for staying fit without leaving your house.",
	"Even with limited space, you can grow a thriving garden. Learn about container gardening, vertical gardens, and best plants for small areas.",
	"Boost your brainpower and reduce anxiety by reading more. Get tips on how to fit reading into your busy schedule and find books you'll love.",
	"Spice up your free time by discovering a new passion. We've compiled a list of engaging hobbies to inspire you and bring joy to your life.",
	"Never run out of books to read again. Curate a personalized reading list that aligns with your interests and expands your knowledge.",
	"Harness the incredible power of positive thinking. Learn how to shift your mindset, overcome challenges, and attract more good into your life.",
	"Take charge of your finances with these straightforward money-saving tips. From budgeting to smart shopping, we'll help you reach your financial goals.",
	"Stay ahead of the curve with our picks for the most innovative tech gadgets. Discover devices that can simplify your life and enhance your productivity.",
}

var tags = []string{
	"sleep",
	"morning-routine",
	"minimalism",
	"productivity",
	"healthy-habits",
	"creativity",
	"travel-hacks",
	"mindfulness",
	"digital-declutter",
	"photography",
	"cooking",
	"sustainability",
	"home-workouts",
	"gardening",
	"reading",
	"hobbies",
	"book-recommendations",
	"positive-thinking",
	"money-saving",
	"tech-gadgets",
}

var comments = []string{
	"This is so helpful! I'm definitely trying the sleep tips tonight.",
	"Great ideas for boosting productivity, especially loved the morning routine section!",
	"Simple living is the way to go. Thanks for the inspiration!",
	"I've been looking for easy healthy habits, and this delivered. Thank you!",
	"Mindfulness in minutes is exactly what I needed. So practical!",
	"As a beginner photographer, these tips are a game-changer.",
	"Love the focus on sustainable living. Every little bit helps!",
	"I never thought about home workouts this way. Time to get moving!",
	"Reading more is always a goal. Great suggestions for a better reading list.",
	"These tech gadgets sound amazing! Adding some to my wishlist.",
}

func Seed(store store.Storage, db *sql.DB) {

	ctx := context.Background()

	users := generateUsers(100)

	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("error seeding user: ", err)
			return
		}
	}

	tx.Commit()
	posts := generatePosts(200, users)

	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("error seeding post: ", err)
			return
		}
	}

	comments := generateComments(500, users, posts)

	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("error seeding comment: ", err)
			return
		}
	}

	log.Println("seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.IntN(len(users))]
		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.IntN(len(titles))],
			Content: contents[rand.IntN(len(contents))],
			Tags: []string{
				tags[rand.IntN(len(tags))],
				tags[rand.IntN(len(tags))],
			},
			Version: 0,
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.IntN(len(posts))].ID,
			UserID:  users[rand.IntN(len(users))].ID,
			Content: contents[rand.IntN(len(comments))],
		}
	}

	return cms
}
