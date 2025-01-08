package todoaggregate

import "time"

type MetaData struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Todo struct {
	ID          string
	Title       string
	Description string
	Status      string
	MetaData
}

func NewTodo(id, title, description, status string) *Todo {
	return &Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
		MetaData:    MetaData{},
	}
}
