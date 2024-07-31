package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	r.HandleFunc("/", app.home)
	r.HandleFunc("/goodgif", app.goodGif)
	r.HandleFunc("/badgif", app.badGif)
	r.HandleFunc("/team/view/{id}", app.teamView)
	r.HandleFunc("/team/gif", app.gifView)
	r.HandleFunc("/team/create", app.teamCreate).Methods("GET")
	r.HandleFunc("/team/create", app.teamCreatePost).Methods("POST")
	r.HandleFunc("/team/vote/{id}", app.teamVote).Methods("GET")
	r.HandleFunc("/team/vote/{id}", app.teamVotePost).Methods("POST")

	return app.recoverPanic(app.logRequest(secureHeaders(r)))
}
