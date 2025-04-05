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
		{
			ID:      3,
			Title:   "Test 3",
			Content: "Test 3",
			Created: time.Now(),
			Expires: time.Now().Add(7 * 24 * time.Hour),
		},
		{
			ID:      4,
			Title:   "Test 4",
			Content: "Test 4",
			Created: time.Now(),
			Expires: time.Now().Add(7 * 24 * time.Hour),
		},
		{
			ID:      5,
			Title:   "Test 5",
			Content: "Test 5",
			Created: time.Now(),
			Expires: time.Now().Add(7 * 24 * time.Hour),
		},
		{
			ID:      6,
			Title:   "Test 6",
			Content: "Test 6",
			Created: time.Now(),
			Expires: time.Now().Add(7 * 24 * time.Hour),
		},
	},
}
