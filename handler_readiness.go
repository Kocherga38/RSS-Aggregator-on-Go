package main

import "net/http"

func handlerReadiness(writer http.ResponseWriter, _ *http.Request) {
	respondWithJSON(writer, 200, struct{}{})
}
