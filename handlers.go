package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/abanseka/gorss/internal/auth"
	"github.com/abanseka/gorss/internal/database"
	"github.com/google/uuid"
)

func handleReady(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}

func handleErr(w http.ResponseWriter, r *http.Request) {
	responswithError(w, 400, "Something went wrong")
}

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

func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request) {
	apikey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		responswithError(w, 403, fmt.Sprintf("Auth error: %v", err))
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apikey)
	if err != nil {
		responswithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
	}

	respondWithJSON(w, 200, dbUserToUser(user))

}
