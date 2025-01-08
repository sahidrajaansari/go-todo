include .env

.PHONY: up down run-app build start
up:
	docker-compose up --build -d --remove-orphans

down:
	docker-compose down

build:
	go build -o ${BINARY} cmd/main.go

clean:
	rm -f ${BINARY}

start: 
	@env DATABASE_USER=${DATABASE_USER} DATABASE_PASSWORD=${DATABASE_PASSWORD} ./${BINARY}

restart: 
	@make clean
	@make build
	@make start