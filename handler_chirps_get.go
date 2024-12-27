package main

import (
	"context"
	"net/http"
	"sort"

	"github.com/google/uuid"
)

type ByOrder []Chirp

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	dbchirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps", err)
		return
	}
	sortDirection := "asc"
	sortDirectionParam := r.URL.Query().Get("sort")
	if sortDirectionParam == "desc" {
		sortDirection = "desc"
	}

	authorID := uuid.Nil
	authorIdString := r.URL.Query().Get("author_id")
	if authorIdString != "" {
		authorID, err = uuid.Parse(authorIdString)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid authorId", err)
			return
		}
	}
	chirps := []Chirp{}
	for _, dbChirp := range dbchirps {
		if authorID != uuid.Nil && dbChirp.UserID != authorID {
			continue
		}

		chirps = append(chirps, Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			UserID:    dbChirp.UserID,
			Body:      dbChirp.Body,
		})
	}

	sort.Slice(chirps, func(i, j int) bool {
		if sortDirection == "desc" {
			return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
		}
		return chirps[i].CreatedAt.Before(chirps[j].CreatedAt)
	})

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
