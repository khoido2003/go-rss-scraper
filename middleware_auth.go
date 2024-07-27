package main

import (
	"fmt"
	"net/http"

	"github.com/khoido2003/go-rss-scraper/internal/auth"
	"github.com/khoido2003/go-rss-scraper/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiconfig) middlewareAuth(handler authHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 401, "Unauthorized")
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintln("Couldn't get user", err))
			return
		}

		handler(w, r, user)
	}

}
