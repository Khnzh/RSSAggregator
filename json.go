package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("500 level error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJson(w http.ResponseWriter, code int, payload any) {
	dat, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Println("Failed to marshal JSON response: ", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
