package main

import (
	"net/http"

	"github.com/harvey-earth/mood/ui"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()
	r.Use(otelmux.Middleware("mood"))

	fileServer := http.FileServer(http.FS(ui.UIFiles))
	r.PathPrefix("/static/").Handler(fileServer)

	r.HandleFunc("/teams/{id}", app.teamView).Methods("GET")
	r.HandleFunc("/teams/{id}/gif", app.gifView).Methods("GET")
	r.HandleFunc("/teams", app.teamCreate).Methods("GET")
	r.HandleFunc("/teams", app.teamCreatePost).Methods("POST")
	r.HandleFunc("/teams/{id}/vote", app.teamVote).Methods("GET")
	r.HandleFunc("/teams/{id}/vote", app.teamVotePost).Methods("POST")
	r.HandleFunc("/ping", ping).Methods("GET")
	r.HandleFunc("/", app.home).Methods("GET")

	return app.recoverPanic(app.logRequest(secureHeaders(r)))
}
