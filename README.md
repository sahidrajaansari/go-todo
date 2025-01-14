# go-todo

## Project Overview

This backend is built in **Go** and structured to follow a clean architecture. Below is a description of the key parts of the project.

### **cmd/server**
This directory contains the entry point and server setup:
- **main.go**: The main entry point of the application.
- **server.go**: Contains the logic for initializing and starting the server.
- **router.go**: The router configuration to handle different routes.
- **todo_routes.go**: Contains the specific routes for the Todo API.

### **config/db**
- **mongo_init.go**: Initializes the connection to MongoDB.

### **pkg/api/handlers**
This directory contains the handler functions that process API requests:
- **handlers.go**: Contains generic handlers for the API.
- **todo_handlers.go**: Contains specific handlers for Todo-related operations.

### **pkg/application/todo**
This directory contains the business logic for the Todo application:
- **mapper.go**: Maps the data between different layers.
- **todo_data_test.go**: Contains test data for the Todo application.
- **todo_service.go**: The business logic for managing Todos.
- **todo_service_test.go**: Contains unit tests for the Todo service.
- **utils.go**: Utility functions used across the application.

### **pkg/contract/todo**
Defines the contract or API interface for Todo operations:
- **create_todo.go**, **get_todo.go**, **update_todo.go**: These files define the structure and logic for creating, retrieving, and updating Todo items.

### **di**
- **wire.go** and **wire_gen.go**: Used for dependency injection and wiring dependencies.

### **domain/persistence**
Contains repositories for managing data:
- **mock_todo_repo.go**: Mock repository for testing.
- **todo_repo.go**: The actual repository interface for data persistence.
- **todo_aggregate**: Handles the Todo domain logic.
- **todo.go**, **todo_data.go**: Represents the Todo aggregate and its data.

### **infrastructure/persistence/todo**
Contains the persistence layer for Todo data:
- **todo_repo.go**: The repository for interacting with MongoDB.
- **todo_data_test.go**: Contains tests for the Todo repository.
- **todo_models.go**: Defines the data models for the Todo application.

### Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/sahidrajaansari/go-todo.git
   cd go-todo
   ```

2. Run Docker containers:
   ```bash
   make up
   ```

3. Run tests:
   ```bash
   make test
   ```

4. Build the Go binary:
   ```bash
   make build
   ```

5. Start the application:
   ```bash
   make start
   ```

6. To test, you can use tools like **Postman** or **cURL** to interact with the Todo API endpoints.

### Running Tests
Run the unit tests:
```bash
make test
```

## Environment Variables

You will need to set up the following environment variables in your `.env` file:

```
DATABASE_NAME=todo-5
DATABASE_USER=admin
DATABASE_PASSWORD=secret
BINARY=todo-5
MONGO_COMPASS_STRING=mongodb://admin:secret@localhost:27017/todo-5?authSource=admin&readPreference=primary&appname=MongDB%20Compass&directConnection=true&ssl=false
```

## Folder Structure

```plaintext
.
├── cmd/
│   └── server/
│       ├── router.go
│       ├── server.go
│       └── todo_routes.go
│   └── main.go
├── config/
│   └── db/
│       └── mongo_init.go
├── pkg/
│   └── api/
│       └── handlers/
│           ├── handlers.go
│           └── todo_handlers.go
│   └── application/
│       └── todo/
│           ├── mapper.go
│           ├── todo_data_test.go
│           ├── todo_service.go
│           ├── todo_service_test.go
│           └── utils.go
│   └── contract/
│       └── todo/
│           ├── create_todo.go
│           ├── get_todo.go
│           └── update_todo.go
├── di/
│   ├── wire.go
│   └── wire_gen.go
├── domain/
│   └── persistence/
│       ├── mock/
│       │   └── mock_todo_repo.go
│       └── todo_repo.go
│   └── todo_aggregate/
│       ├── todo.go
│       └── todo_data.go
├── infrastructure/
│   └── persistence/
│       └── todo/
│           ├── todo_data_test.go
│           ├── todo_models.go
│           ├── todo_repo.go
│           └── todo_repo_test.go
├── script/
│   ├── post_all_todo.sh
│   └── go-todo.postman_collection.json
└── .env
└── .gitignore
└── Makefile
└── README.md
└── docker-compose.yml
└── go.mod
└── go.sum
```

