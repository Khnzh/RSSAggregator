package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Khnzh/RSSAggregator/internal/database"
	"github.com/google/uuid"
)

func (api *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:", err))
		return
	}

	feed, err := api.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedSerializer(feed))
}

func (api *apiConfig) handleFetchUserFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := api.DB.FetchFeedsByUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedsSerializer(feeds))
}
