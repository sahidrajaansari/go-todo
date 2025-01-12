package todo_aggregate

import "time"

var TodoAgg = Todo{
	ID:          "2rNWzlOcXbhQAkRRk77StyDYAqf",
	Title:       "Buy groceries",
	Description: "Milk, bread, eggs, fruits",
	Status:      "pending",
	MetaData: MetaData{
		CreatedAt: time.Date(2025, 1, 10, 4, 52, 1, 0, time.UTC), // Fixed timestamp
		UpdatedAt: time.Date(2025, 1, 12, 4, 52, 1, 0, time.UTC), // Fixed timestamp
	},
}
