package models

type MockSnippetModel struct{}

func (m *MockSnippetModel) Insert(title, content string, expires int) (int, error) {
	return 124, nil
}

func (m *MockSnippetModel) Get(id int) (Snippet, error) {
	return Snippet{
		ID:    1,
		Title: "Mock Title",
	}, nil
}
