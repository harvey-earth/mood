package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	r.HandleFunc("/team/{id}/view", app.teamView).Methods("GET")
	r.HandleFunc("/team/{id}/gif", app.gifView).Methods("GET")
	r.HandleFunc("/team/create", app.teamCreate).Methods("GET")
	r.HandleFunc("/team/create", app.teamCreatePost).Methods("POST")
	r.HandleFunc("/team/{id}/vote", app.teamVote).Methods("GET")
	r.HandleFunc("/team/{id}/vote", app.teamVotePost).Methods("POST")
	r.HandleFunc("/", app.home).Methods("GET")

	return app.recoverPanic(app.logRequest(secureHeaders(r)))
}
