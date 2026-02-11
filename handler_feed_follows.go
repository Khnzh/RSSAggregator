package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Khnzh/RSSAggregator/internal/database"
	"github.com/google/uuid"
)

func (api *apiConfig) handleCreateFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:", err))
		return
	}

	follow, err := api.DB.CreateFollow(r.Context(), database.CreateFollowParams{
		ID:     uuid.New(),
		FeedID: params.FeedID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create a follow: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFollowSerializer(follow))
}

func (api *apiConfig) handleFetchUserFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	follows, err := api.DB.FetchFollowsByUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFollowsSerializer(follows))
}

func (api *apiConfig) handleFetchUserFollowedFeeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := api.DB.FetchFeedsFollowedByUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedsSerializer(feeds))
}
