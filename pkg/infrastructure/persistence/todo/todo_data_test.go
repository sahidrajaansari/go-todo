package todo

import (
	"go.mongodb.org/mongo-driver/bson"
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

// Correct MongoDB data package 3: Represents a task to go for a run
var correctMongoData3 = bson.D{
	{Key: "id", Value: "3zUY2V8l1gN3opJk72Vr5uXjT1F"},
	{Key: "title", Value: "Go for a Run"},
	{Key: "description", Value: "Run 5 kilometers in the park"},
	{Key: "status", Value: "Completed"},
	{Key: "createdat", Value: "2025-01-12T07:45:00Z"},
	{Key: "updatedat", Value: "2025-01-12T08:30:00Z"},
}

// Correct MongoDB data package 4: Represents a task to read a book
var correctMongoData4 = bson.D{
	{Key: "id", Value: "1nBV5Z0OqZpQ2Kw3HcMd4Y5W4Tp"},
	{Key: "title", Value: "Read a Book"},
	{Key: "description", Value: "Read 30 pages of a novel"},
	{Key: "status", Value: "Pending"},
	{Key: "createdat", Value: "2025-01-08T20:00:00Z"},
	{Key: "updatedat", Value: ""},
}
