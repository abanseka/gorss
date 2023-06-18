package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/abanseka/gorss/internal/database"
	"github.com/google/uuid"
)

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
	respondWithJSON(w, 200, dbUserToUser(user))
}

func handleReady(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}

func handleErr(w http.ResponseWriter, r *http.Request) {
	responswithError(w, 400, "Something went wrong")
}
