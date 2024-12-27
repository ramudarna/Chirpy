package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/ramudarna/Chirpy/internal/auth"
)

func (cfg *apiConfig) handlerChirpDeleteById(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	pathValue := r.PathValue("chirpID")
	chirpId, err := uuid.Parse(pathValue)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid chirp Id", err)
		return
	}

	chirp, err := cfg.db.GetChirp(context.Background(), chirpId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "couldn't get chirp", err)
		return
	}

	if chirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "user is not author", errors.New("forbidden for user"))
		return
	}

	err = cfg.db.DeleteChirpById(context.Background(), chirpId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't find chirp", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
