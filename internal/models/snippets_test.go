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

func TestInsertSnippet(t *testing.T) {
	db, err := helpers.SetupTestDB(t)

	if err != nil {
		t.Fatal(err)
	}

	err = helpers.LoadTestDataFromStruct(t, db, testdata.DBFake1)
	if err != nil {
		t.Fatal(err)
	}

	model := &SnippetModel{DB: db}
	idSnippet, err := model.Insert("title1", "content1", 7)
	if err != nil {
		t.Fatalf("expected to find snippet: %v", err)
	}

	snippet, err := model.Get(idSnippet)
	if err != nil {
		t.Fatalf("expected to find snippet: %v", err)
	}

	if snippet.Title != "title1" {
		t.Errorf("expected title 'title1'; got %q", snippet.Title)
	}

}

func TestGetLatest(t *testing.T) {
	db, err := helpers.SetupTestDB(t)

	if err != nil {
		t.Fatal(err)
	}

	err = helpers.LoadTestDataFromStruct(t, db, testdata.DBFake1)
	if err != nil {
		t.Fatal(err)
	}

	model := &SnippetModel{DB: db}
	snippet, err := model.Latest()
	if err != nil {
		t.Fatalf("expected to find snippet: %v", err)
	}

	if len(snippet) == 0 {
		t.Fatalf("expected to find at least one snippet; got %d", len(snippet))
	}
}
