package todo

import (
	"errors"
	"net/url"
	"time"
	tContracts "todo-level-5/pkg/contract/todo"
	"todo-level-5/pkg/domain/todo_aggregate"
)

// HTTP method constants
const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

// createGinContextHelper struct to hold method, URL, and parameters
type createGinContextHelper struct {
	Method string
	URL    string
	Params map[string]string
}

// Example instance of createGinContextHelper
var createGinContextHelper1 = createGinContextHelper{
	Method: GET,
	URL:    "/todos",
	Params: map[string]string{
		"status": "pending",
		"user":   "123",
	},
}

// getUrl method to return the full URL
func (c *createGinContextHelper) getUrl() string {
	// Construct the base URL
	baseURL := c.URL
	if len(c.Params) > 0 {
		query := c.getRawQuery()
		baseURL = baseURL + "?" + query
	}
	return baseURL
}

// getRawQuery method to return the raw query string
func (c *createGinContextHelper) getRawQuery() string {
	// Construct the raw query string
	if len(c.Params) > 0 {
		query := url.Values{}
		for key, value := range c.Params {
			query.Add(key, value)
		}
		// Return the raw query string without the base URL
		return query.Encode()
	}
	return ""
}

// Errors
var (
	errInvalidTodoID   = errors.New("invalid todo ID")
	errRepositoryError = errors.New("database error: unable to insert new todo")
	validTodoID        = "2rNWzlOcXbhQAkRRk77StyDYAqf"
	nonExistentID      = "121212311"
)

// Prototype data for CreateTodoRequest
var ValidCreateTodoRequest = tContracts.CreateTodoRequest{
	Title:       "Learn GoLang",
	Description: "Complete the GoLang tutorials on concurrency and testing.",
	Status:      "pending",
}

var ValidUpdateTodoRequestAllFields = tContracts.UpdateTodoRequest{
	Title:       "Complete GoLang tutorial",
	Description: "Finish all exercises related to GoLang concurrency.",
	Status:      "in-progress",
}

var ValidUpdateTodoRequestSingleField = tContracts.UpdateTodoRequest{
	Title:       "Complete GoLang tutorial",
	Description: "",
	Status:      "",
}

// Function to generate a sample Todo object
func CreateSampleTodo(id, title, description, status string, createdAt, updatedAt time.Time) todo_aggregate.Todo {
	return todo_aggregate.Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
		MetaData: todo_aggregate.MetaData{
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}
}

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

	// todoAgg3 = CreateSampleTodo(
	// 	"5uOWzlHfXbhQAkRRk00StyGWAe",
	// 	"Clean the house",
	// 	"Living room, kitchen, and bedrooms",
	// 	"pending",
	// 	time.Date(2025, 1, 8, 9, 15, 0, 0, time.UTC),
	// 	time.Date(2025, 1, 8, 9, 15, 0, 0, time.UTC),
	// )
)
