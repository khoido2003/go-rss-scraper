package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payoad interface{}) {
	dat, err := json.Marshal(payoad)

	if err != nil {
		log.Printf("Failed to marshal JSON response %v", payoad)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
