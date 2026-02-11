package main

import (
	"time"

	"github.com/Khnzh/RSSAggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID     uuid.UUID `json:"id"`
	FeedID uuid.UUID `json:"feed_id"`
	UserID uuid.UUID `json:"user_id"`
}

func databaseUserSerializer(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

func databaseFeedSerializer(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func databaseFeedsSerializer(feeds []database.Feed) []Feed {
	result := []Feed{}
	for _, feed := range feeds {
		result = append(result, databaseFeedSerializer(feed))
	}
	return result
}

func databaseFollowSerializer(follow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:     follow.ID,
		UserID: follow.UserID,
		FeedID: follow.FeedID,
	}
}

func databaseFollowsSerializer(follows []database.FeedFollow) []FeedFollow {
	result := []FeedFollow{}
	for _, follow := range follows {
		result = append(result, databaseFollowSerializer(follow))
	}
	return result
}
