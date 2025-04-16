package main

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"snippetbox.tonidefez.net/internal/models"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func TestPing(t *testing.T) {
	rr := httptest.NewRecorder()

	r := httptest.NewRequest("GET", "/ping", nil)

	// app := newTestApplication() // Lo vamos a definir luego si hace falta
	handler := http.HandlerFunc(pingHandler)

	handler.ServeHTTP(rr, r)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200; got %d", rr.Code)
	}

	if rr.Body.String() != "OK" {
		t.Errorf("expected body 'OK'; got %q", rr.Body.String())
	}
}

func TestHome(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	// simulate dependencies
	dummyLogger := slog.New(slog.NewTextHandler(io.Discard, nil))
	templateCache, err := newTemplateCache()

	if err != nil {
		t.Fatal(err)
	}

	app := &Application{
		logger:        dummyLogger,
		snippets:      &models.MockSnippetModel{},
		templateCache: templateCache,
	}

	handler := http.HandlerFunc(app.Home)

	handler.ServeHTTP(rr, req)

	// verify status
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200; got %d", rr.Code)
	}

	// verify header
	got := rr.Header().Get("Server")
	if got != "Go" {
		t.Errorf("expected header 'Server: Go'; got %q", got)
	}

	// verify content HTML (optional)
	if !strings.Contains(rr.Body.String(), "<title>Home - Snippetbox</title>") {
		t.Errorf("expected HTML title; got %q", rr.Body.String())
	}
}

func TestSnippetView(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/snippet/view/1", nil)

	// simulate dependencies
	dummyLogger := slog.New(slog.NewTextHandler(io.Discard, nil))
	templateCache, err := newTemplateCache()
	if err != nil {
		t.Fatal(err)
	}

	app := &Application{
		logger:        dummyLogger,
		snippets:      &models.MockSnippetModel{},
		templateCache: templateCache,
	}

	router := app.routes()
	router.ServeHTTP(rr, req)

	// verify status
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200; got %d", rr.Code)
	}

	body := rr.Body.String()

	if !strings.Contains(body, "Mock Title") {
		t.Errorf("expected body to contain 'Mock Title'; got %q", body)
	}

}

func TestSnippetCreateGet(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/snippet/create", nil)

	// simulate dependencies
	dummyLogger := slog.New(slog.NewTextHandler(io.Discard, nil))
	dummyDB := &models.SnippetModel{DB: nil}
	templateCache, err := newTemplateCache()

	if err != nil {
		t.Fatal(err)
	}

	app := &Application{
		logger:        dummyLogger,
		snippets:      dummyDB,
		templateCache: templateCache,
	}

	router := app.routes()
	router.ServeHTTP(rr, req)

	// verify status
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200; got %d", rr.Code)
	}
}

func TestSnippetCreatePost(t *testing.T) {
	rr := httptest.NewRecorder()

	form := "title=Mi+Título&content=Texto+del+snippet&expires=7"
	req := httptest.NewRequest("POST", "/snippet/create", strings.NewReader(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	dummyLogger := slog.New(slog.NewTextHandler(io.Discard, nil))

	// simulate dependencies
	app := &Application{
		logger:   dummyLogger,
		snippets: &models.MockSnippetModel{},
	}

	router := app.routes()
	router.ServeHTTP(rr, req)

	// Verificar redirección
	if rr.Code != http.StatusSeeOther {
		t.Errorf("expected status 303 See Other; got %d", rr.Code)
	}

	// Verificar Location
	expectedLocation := "/snippet/view/124"
	actualLocation := rr.Header().Get("Location")
	if actualLocation != expectedLocation {
		t.Errorf("expected Location header %q; got %q", expectedLocation, actualLocation)
	}
}

func TestSnippetCreatePost_InvalidData(t *testing.T) {
	templateCache, err := newTemplateCache()

	if err != nil {
		t.Fatal(err)
	}
	app := &Application{
		templateCache: templateCache,
	}

	form := url.Values{}
	form.Add("title", "")
	form.Add("content", "Some valid content")
	form.Add("expires", "7")

	req := httptest.NewRequest(http.MethodPost, "/snippet/create", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	app.SnippetCreatePost(rr, req)

	res := rr.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("expected status 422 OK; got %d", res.StatusCode)
	}
}
