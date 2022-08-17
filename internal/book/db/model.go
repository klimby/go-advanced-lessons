package db

import (
	"database/sql"
	"github.com/klimby/go-advanced-lessons/internal/author"
	"github.com/klimby/go-advanced-lessons/internal/book"
)

type Book struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Age     sql.NullInt32   `json:"age"`
	Authors []author.Author `json:"authors"`
}

func (m *Book) ToDomain() book.Book {
	b := book.Book{
		ID:   m.ID,
		Name: m.Name,
	}
	if m.Age.Valid {
		b.Age = int(m.Age.Int32)
	}

	return b
}
