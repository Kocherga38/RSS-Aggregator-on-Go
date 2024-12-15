package main

import (
	"fmt"
	"net/http"

	"github.com/Kocherga38/rssagg/internal/auth"
	"github.com/Kocherga38/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (APICfg *APIConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		apiKey, err := auth.GetAPIKey(request.Header)
		if err != nil {
			respondWithError(writer, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := APICfg.DB.GetUserByAPIKey(request.Context(), apiKey)
		if err != nil {
			respondWithError(writer, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(writer, request, user)
	}
}
