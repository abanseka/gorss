package main

import "net/http"

func handleReady(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}

func handleErr(w http.ResponseWriter, r *http.Request) {
	responswithError(w, 400, "Something went wrong")
}
