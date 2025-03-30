package models

import (
	"testing"

	"snippetbox.tonidefez.net/internal/helpers"
	"snippetbox.tonidefez.net/internal/testdata"
)

func TestGetSnippet_WithLoadedData(t *testing.T) {
	db, err := helpers.SetupTestDB(t)

	if err != nil {
		t.Fatal(err)
	}

	err = helpers.LoadTestDataFromStruct(t, db, testdata.DBFake1)
	if err != nil {
		t.Fatal(err)
	}

	model := &SnippetModel{DB: db}
	snippet, err := model.Get(1)
	if err != nil {
		t.Fatalf("expected to find snippet: %v", err)
	}

	if snippet.Title != "Primer snippet" {
		t.Errorf("expected title 'Primer snippet'; got %q", snippet.Title)
	}
}
