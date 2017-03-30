package main

import (
	"net/http"
	"encoding/json"
)

type ResponseError struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, message string) {

	resp := ResponseError{
		Error:message,
	}

	respondWithJSON(w, code, resp)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
