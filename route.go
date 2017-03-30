package main

import (
	"net/http"
)

type OSResponse struct {
	OS string `json:"os"`
}

func (a *App) serveOs(w http.ResponseWriter, r *http.Request) {
	hostOS, err := NewInfo().Host()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := OSResponse{}
	resp.OS = hostOS

	respondWithJSON(w, http.StatusOK, resp)
}

type MemoryResponse struct {
	Total uint64  `json:"total"`
	Free  uint64  `json:"free"`
	Used  float64 `json:"used"`
}

func (a *App) serveMemory(w http.ResponseWriter, r *http.Request) {
	var resp MemoryResponse
	var err error
	resp.Total, resp.Free, resp.Used, err = NewInfo().Memory()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, resp)
}
