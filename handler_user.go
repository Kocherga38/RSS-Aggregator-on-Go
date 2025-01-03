package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Kocherga38/rssagg/internal/database"
	"github.com/google/uuid"
)

func (APICfg *APIConfig) handlerCreateUser(writer http.ResponseWriter, request *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(request.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal("Couldn't load location")
	}

	nowMoscowTime := time.Now().In(location)

	user, err := APICfg.DB.CreateUser(request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: nowMoscowTime,
		UpdatedAt: nowMoscowTime,
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(writer, 201, databaseUserToUser(user))
}

func (APICfg *APIConfig) handlerGetUser(writer http.ResponseWriter, request *http.Request, user database.User) {
	respondWithJSON(writer, 200, databaseUserToUser(user))
}

func (APICfg *APIConfig) handlerGetPostsForUser(writer http.ResponseWriter, request *http.Request, user database.User) {
	posts, err := APICfg.DB.GetPostsForUser(request.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	respondWithJSON(writer, 200, databasePostsToPosts(posts))
}
