package main

import (
	"errors"
	"net/http"
	"os"
)

func (cfg *apiConfig) handleReset(w http.ResponseWriter, r *http.Request) {
	platform := os.Getenv("PLATFORM")
	if platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Forbidden", errors.New("forbidden request"))
		return
	}
	ctx := r.Context()

	err := cfg.db.ResetUsers(ctx)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not reset users", err)
	}

	cfg.fileserverHits.Store(0)
	respondWithJSON(w, http.StatusOK, []byte("Counter reset"))

}
