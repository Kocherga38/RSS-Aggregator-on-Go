package main

import "net/http"

func handlerError(writer http.ResponseWriter, request *http.Request) {
	respondWithError(writer, 400, "Something went wrong")
}
