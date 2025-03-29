package models

type SnippetModeler interface {
	Insert(title, content string, expires int) (int, error)
}
