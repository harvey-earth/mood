package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

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

func (app *application) badGif(w http.ResponseWriter, r *http.Request) {
	lissajous(w, 100, palette2)
}

func (app *application) goodGif(w http.ResponseWriter, r *http.Request) {
	lissajous(w, 2, palette1)
}

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

func (app *application) teamCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	name := "Team 1"

	id, err := app.teams.Insert(name)
	if err != nil {
		app.serverError(w, err)
	}
	http.Redirect(w, r, fmt.Sprintf("/team/view?id=%d", id), http.StatusSeeOther)
}

func (app *application) teamView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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
	// lissajous(w, float64(team.Score), palette1)
	// fmt.Fprintf(w, "%+v", team)
}
