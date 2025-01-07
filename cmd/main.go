package main

import (
	"context"
	"todo-level-5/config/db"
)

func main() {
	db.Connect(context.Background())
}
