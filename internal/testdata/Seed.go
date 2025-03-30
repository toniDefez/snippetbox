package testdata

import (
	"time"
)

// same structure that models.Snippet but just for testing
type SnippetFake struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type Seed struct {
	Snippets []SnippetFake
}

var DBFake1 = Seed{
	Snippets: []SnippetFake{
		{
			ID:      1,
			Title:   "Primer snippet",
			Content: "Test 1",
			Created: time.Now(),
			Expires: time.Now().Add(7 * 24 * time.Hour),
		},
		{
			ID:      2,
			Title:   "Test 2",
			Content: "Test 2",
			Created: time.Now(),
			Expires: time.Now().Add(7 * 24 * time.Hour),
		},
	},
}
