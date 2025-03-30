package models

import "testing"

func TestGetSnippet_WithLoadedData(t *testing.T) {
	startMySQLContainer(t)

	dsn := "testuser:testpass@tcp(127.0.0.1:3307)/testdb?parseTime=true"
	db := waitForMySQLReady(t, dsn)

	_, err := db.Exec(`CREATE TABLE snippets (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created DATETIME NOT NULL,
		expires DATETIME NOT NULL
	);`)
	if err != nil {
		t.Fatal(err)
	}

	loadTestDataFromJSON(t, db, "../../testdata/snippets.json")

	model := &SnippetModel{DB: db}
	snippet, err := model.Get(1)
	if err != nil {
		t.Fatalf("expected to find snippet: %v", err)
	}

	if snippet.Title != "Primer snippet" {
		t.Errorf("expected title 'Primer snippet'; got %q", snippet.Title)
	}
}
