# Include environment variables from the .env file
include .env

# Define targets as phony to prevent conflicts with file names
.PHONY: up down run-app build start clean restart

# Bring up the Docker containers, build images if necessary, and remove any orphaned containers
up:
	docker-compose up --build -d --remove-orphans

# Shut down the Docker containers and clean up resources
down:
	docker-compose down

# Build the Go binary from the main Go file
build:
	go build -o ${BINARY} cmd/main.go

# Remove the compiled binary
clean:
	rm -f ${BINARY}

# Start the application using environment variables for database credentials
start: 
	@env DATABASE_USER=${DATABASE_USER} DATABASE_PASSWORD=${DATABASE_PASSWORD} ./${BINARY}

test:
	go test ./... -v -cover

# Restart the application by cleaning, rebuilding, and starting it
restart: clean build start

