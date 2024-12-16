package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Kocherga38/rssagg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (APICfg *APIConfig) handlerCreateFeedFollow(writer http.ResponseWriter, request *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
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

	feedFollow, err := APICfg.DB.CreateFeedFollow(request.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: nowMoscowTime,
		UpdatedAt: nowMoscowTime,
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	respondWithJSON(writer, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (APICfg *APIConfig) handlerGetFeedFollows(writer http.ResponseWriter, request *http.Request, user database.User) {
	feedFollows, err := APICfg.DB.GetFeedFollow(request.Context(), user.ID)
	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}

	respondWithJSON(writer, 201, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (APICfg *APIConfig) handlerDeleeteFeedFollows(writer http.ResponseWriter, request *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(request, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Couldn't parse feed follow id: %v", err))
		return
	}

	err = APICfg.DB.DeleteFeedFollow(request.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}

	respondWithJSON(writer, 200, struct{}{})
}
