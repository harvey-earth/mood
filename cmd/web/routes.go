package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/goodgif", app.goodGif)
	mux.HandleFunc("/badgif", app.badGif)
	mux.HandleFunc("/team/view", app.teamView)
	mux.HandleFunc("/team/gif", app.gifView)
	mux.HandleFunc("/team/create", app.teamCreate)
	// mux.HandleFunc("/team/vote", app.teamVote)

	return mux
}
