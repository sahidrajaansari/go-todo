package main

import (
	"todo-level-5/cmd/server"
)

func main() {
	// := db.Connect(context.Background())

	srv := server.NewHttpServer()
	server.SetupRoutes(srv)
	srv.Run(":8080")
}
