package todo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	validTodoID       = "Valid"
	nonExistentTodoID = "Non"
	// inValidTodoID     = "invaild"
)

// Can be Used To Generate a Bson Error
var errorMongoData = bson.D{
	{Key: "id", Value: "2rUJjjM3txRh6prtuwbPrl399tp"},
	{Key: "title", Value: "Buy Groceries"},
	{Key: "description", Value: "Include fruits and vegetables"},
	{Key: "status", Value: true},
}

// Correct MongoDB data package 1
var correctMongoData1 = bson.D{
	{Key: "id", Value: "2rUJjjM3txRh6prtuwbPrl399tp"},
	{Key: "title", Value: "Buy Groceries"},
	{Key: "description", Value: "Include fruits and vegetables"},
	{Key: "status", Value: "Pending"},
	{Key: "createdat", Value: ""}, // No creation timestamp yet
	{Key: "updatedat", Value: ""}, // No update timestamp yet
}

// Correct MongoDB data package 2: Represents a task to complete homework
var correctMongoData2 = bson.D{
	{Key: "id", Value: "5gHJ8nK1jlM7qRkiGgNz1tRt9Vb"},
	{Key: "title", Value: "Complete Homework"},
	{Key: "description", Value: "Finish math and science assignments"},
	{Key: "status", Value: "In Progress"},
	{Key: "createdat", Value: "2025-01-10T10:00:00Z"},
	{Key: "updatedat", Value: "2025-01-11T15:30:00Z"},
}

var (
	validQuery   = "status=pending"
	invaildQuery = "status=invalid%query"
)

// Prototype Todos
var (
	todoAgg1 = CreateSampleTodo(
		"3sMWzlYdXbhQAkRRk88StyEYBr",
		"Prepare presentation",
		"Slides for the quarterly review meeting",
		"in-progress",
		time.Date(2025, 1, 9, 10, 30, 0, 0, time.UTC),
		time.Date(2025, 1, 11, 14, 20, 0, 0, time.UTC),
	)

	todoAgg2 = CreateSampleTodo(
		"4tNXzlFdXbhQAkRRk99StyFZCs",
		"Read book",
		"Complete 'Atomic Habits' by James Clear",
		"completed",
		time.Date(2025, 1, 5, 8, 45, 0, 0, time.UTC),
		time.Date(2025, 1, 6, 12, 15, 0, 0, time.UTC),
	)
)
