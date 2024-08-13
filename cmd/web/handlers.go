package main

import (
	"errors"
	"fmt"
	"html/template"
	"image/color"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/harvey-earth/mood/internal/models"
	"github.com/harvey-earth/mood/ui"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"html/base.tmpl.html",
		"html/partials/nav.tmpl.html",
		"html/pages/home.tmpl.html",
	}
	ts, err := template.ParseFS(ui.UIFiles, files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

// Takes team id and gets score from database. Returns a lissajous gif based on score.
func (app *application) gifView(w http.ResponseWriter, r *http.Request) {
	// Color patterns for gifs. c1 is green and c5 is red
	var c1 = []color.Color{color.RGBA{98, 203, 49, 255}, color.Black}
	var c2 = []color.Color{color.RGBA{255, 182, 6, 255}}
	var c3 = []color.Color{color.RGBA{230, 126, 34, 255}}
	var c4 = []color.Color{color.RGBA{231, 76, 60, 255}}
	var c5 = []color.Color{color.RGBA{255, 0, 0, 255}, color.Black}
	// color will hold the palette color of the gif
	var color []color.Color

	// Get team by ID
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
	currentScore := team.Score

	// Set color palette based on score
	switch {
	case currentScore >= 1 && currentScore <= 20:
		color = c1
	case currentScore > 20 && currentScore <= 40:
		color = c2
	case currentScore > 40 && currentScore <= 60:
		color = c3
	case currentScore > 60 && currentScore <= 80:
		color = c4
	default:
		color = c5
	}

	// Return gif
	lissajous(w, float64(team.Score), color)
}

// Returns form to create a new team
func (app *application) teamCreate(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"html/base.tmpl.html",
		"html/partials/nav.tmpl.html",
		"html/pages/create.tmpl.html",
	}

	ts, err := template.ParseFS(ui.UIFiles, files...)
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
	http.Redirect(w, r, fmt.Sprintf("/team/%d/view", id), http.StatusSeeOther)
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
		"html/base.tmpl.html",
		"html/partials/nav.tmpl.html",
		"html/pages/view.tmpl.html",
	}
	ts, err := template.ParseFS(ui.UIFiles, files...)
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
		"html/base.tmpl.html",
		"html/partials/nav.tmpl.html",
		"html/pages/vote.tmpl.html",
	}
	ts, err := template.ParseFS(ui.UIFiles, files...)
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
	http.Redirect(w, r, fmt.Sprintf("/team/%d/view", id), http.StatusSeeOther)
}
