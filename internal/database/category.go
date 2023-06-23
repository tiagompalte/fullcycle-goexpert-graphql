package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	sts := "CREATE TABLE IF NOT EXISTS categories(id string, name string NOT NULL, description string);"
	_, err := db.Exec(sts)
	if err != nil {
		panic(err)
	}
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories(id,name,description) VALUES ($1,$2,$3)", id, name, description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}
