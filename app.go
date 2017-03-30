package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()

	a.SetRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) SetRoutes() {
	a.Router.HandleFunc("/memory",
		a.serveMemory).Methods("GET")
	a.Router.HandleFunc("/os",
		a.serveOs).Methods("GET")
}
