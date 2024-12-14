package main

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	dbchirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps", err)
		return
	}
	chirps := []Chirp{}
	for _, chirp := range dbchirps {
		chirps = append(chirps, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}
	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handleChirpsGetById(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("chirpID")
	id, err := uuid.Parse(pathValue)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid chirp Id", err)
		return
	}
	chirp, err := cfg.db.GetChirp(context.Background(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "couldn't get chirp", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		UserID:    chirp.UserID,
		Body:      chirp.Body,
	})
}
