package server

import (
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, "Hello, World")
}

func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "Internal Server Error")
}
