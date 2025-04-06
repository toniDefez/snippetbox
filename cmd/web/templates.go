package main

import (
	"path/filepath"
	"runtime"
	"text/template"

	"snippetbox.tonidefez.net/internal/models"
)

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
// At the moment it only contains one field, but we'll add more
// to it as the build progresses.
type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	_, filename, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(filename), "..", "..", "ui", "html")

	pages, err := filepath.Glob(filepath.Join(basePath, "pages", "*.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(filepath.Join(basePath, "base.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(basePath, "partials", "*.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
