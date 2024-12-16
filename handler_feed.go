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

func (APICfg *APIConfig) handlerCreateFeed(writer http.ResponseWriter, request *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
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

	feed, err := APICfg.DB.CreateFeed(request.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: nowMoscowTime,
		UpdatedAt: nowMoscowTime,
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	respondWithJSON(writer, 201, databaseFeedToFeed(feed))
}

func (APICfg *APIConfig) handlerGetFeeds(writer http.ResponseWriter, request *http.Request) {
	feeds, err := APICfg.DB.GetFeeds(request.Context())

	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	respondWithJSON(writer, 201, databaseFeedsToFeeds(feeds))
}
