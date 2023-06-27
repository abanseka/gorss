package main

import (
	"fmt"
	"net/http"

	"github.com/abanseka/gorss/internal/auth"
	"github.com/abanseka/gorss/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responswithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			responswithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
		}

		handler(w, r, user)
	}
}
