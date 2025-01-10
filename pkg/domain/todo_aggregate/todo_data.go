package todo_aggregate

import "time"

var TodoAgg = Todo{
	ID:          "1",
	Title:       "Buy groceries",
	Description: "Milk, bread, eggs, fruits",
	Status:      "pending",
	MetaData: MetaData{
		CreatedAt: time.Now().Add(-48 * time.Hour), // Created 2 days ago
		UpdatedAt: time.Now().Add(-24 * time.Hour), // Updated 1 day ago
	},
}
