package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/marktlinn/RssScrape/auth"
	"github.com/marktlinn/RssScrape/internal/database"
	"github.com/marktlinn/RssScrape/models"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, "Hello, World")
}

func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "Internal Server Error")
}

func (s *ServerConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	p := params{}

	if err := decoder.Decode(&p); err != nil && err != io.EOF {
		respondWithError(w, 400, fmt.Sprintf("failed to decode params: %s\n", err))
	}

	user, err := s.DB.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      p.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("failed to create User with params %s in DB: %s\n", p.Name, err))
		return
	}

	respondWithJson(w, 200, models.DatabaseUserToUser(user))
}

func (s *ServerConfig) handlerGetUserByID(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("unauthorized: %s\n", err))
		return
	}

	user, err := s.DB.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("failed to find User: %s\n", err))
		return
	}

	respondWithJson(w, 200, models.DatabaseUserToUser(user))
}
