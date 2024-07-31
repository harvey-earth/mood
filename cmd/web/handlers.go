package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/harvey-earth/mood/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

// This function needs to be converted to a static asset to serve from homepage
func (app *application) badGif(w http.ResponseWriter, r *http.Request) {
	lissajous(w, 50, palette2)
}

// This function needs to be converted to a static asset to serve from homepage
func (app *application) goodGif(w http.ResponseWriter, r *http.Request) {
	lissajous(w, 1, palette1)
}

// Takes team id and gets score from database. Returns a lissajous gif based on score.
func (app *application) gifView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.notFound(w)
		return
	}
	team, err := app.teams.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	lissajous(w, float64(team.Score), palette1)
}

// Returns form to create a new team
func (app *application) teamCreate(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/create.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

// Creates new team and redirects to view
func (app *application) teamCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	name := r.PostForm.Get("team-name")

	id, err := app.teams.Insert(name)
	if err != nil {
		app.serverError(w, err)
	}
	http.Redirect(w, r, fmt.Sprintf("/team/view/%d", id), http.StatusSeeOther)
}

// Returns page showing team information
func (app *application) teamView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 0 {
		app.notFound(w)
		return
	}

	team, err := app.teams.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/view.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", team)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) teamVote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	team, err := app.teams.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/vote.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", team)
	if err != nil {
		app.serverError(w, err)
	}
}
func (app *application) teamVotePost(w http.ResponseWriter, r *http.Request) {
	var newScore int
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	team, err := app.teams.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	err = r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	switch value := r.PostForm.Get("vote"); value {
	case "5":
		newScore = max(team.Score-5, 2)
	case "4":
		newScore = max(team.Score-2, 2)
	case "3":
		newScore = team.Score
	case "2":
		newScore = max(team.Score+2, 2)
	case "1":
		newScore = max(team.Score+5, 2)
	default:
		// Error condition
		app.clientError(w, http.StatusBadRequest)

	}
	err = app.teams.Update(id, newScore)
	if err != nil {
		app.serverError(w, err)
	}
	http.Redirect(w, r, fmt.Sprintf("/team/view/%d", id), http.StatusSeeOther)
}
