package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
	"text/template"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	_, filename, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(filename), "..", "..", "ui", "html")

	files := []string{
		filepath.Join(basePath, "base.tmpl"),
		filepath.Join(basePath, "partials", "nav.tmpl"),
		filepath.Join(basePath, "pages", "home.tmpl"),
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *Application) SnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *Application) SnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func (app *Application) SnippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
