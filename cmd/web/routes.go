package main

import (
	"net/http"

	"github.com/harvey-earth/mood/ui"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	fileServer := http.FileServer(http.FS(ui.UIFiles))
	r.PathPrefix("/static/").Handler(fileServer)

	r.HandleFunc("/team/{id}/view", app.teamView).Methods("GET")
	r.HandleFunc("/team/{id}/gif", app.gifView).Methods("GET")
	r.HandleFunc("/team/create", app.teamCreate).Methods("GET")
	r.HandleFunc("/team/create", app.teamCreatePost).Methods("POST")
	r.HandleFunc("/team/{id}/vote", app.teamVote).Methods("GET")
	r.HandleFunc("/team/{id}/vote", app.teamVotePost).Methods("POST")
	r.HandleFunc("/ping", ping).Methods("GET")
	r.HandleFunc("/", app.home).Methods("GET")

	return app.recoverPanic(app.logRequest(secureHeaders(r)))
}
