package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/abanseka/gorss/internal/database"
	"github.com/google/uuid"
)

// active server
func handleReady(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}

// handle errors
func handleErr(w http.ResponseWriter, r *http.Request) {
	responswithError(w, 400, "Something went wrong")
}

// handle user creation
func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responswithError(w, 400, fmt.Sprintf("Error parsin JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		responswithError(w, 400, fmt.Sprintf("Couldn't create user: %s", err))
		return
	}

	respondWithJSON(w, 201, dbUserToUser(user))
}

// handle user retrieval
func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, dbUserToUser(user))
}

// handle feed creation
func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responswithError(w, 400, fmt.Sprintf("Error parsin JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		responswithError(w, 400, fmt.Sprintf("Couldn't create feed: %s", err))
		return
	}

	respondWithJSON(w, 201, dbFeedToFeed(feed))
}

// handle feeds retrieval
func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feed, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		responswithError(w, 400, fmt.Sprintf("Couldn't get feed: %s", err))
		return
	}

	respondWithJSON(w, 201, dbFeedsToFeeds(feed))
}

// handle feed creation
func (apiCfg *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responswithError(w, 400, fmt.Sprintf("Error parsin JSON: %v", err))
		return
	}

	FeedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		responswithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %s", err))
		return
	}

	respondWithJSON(w, 201, dbFeedFollowToFollow(FeedFollow))
}
