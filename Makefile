.PHONY: start stop run-app

start:
	docker-compose up --build -d --remove-orphans

stop:
	docker-compose down

run-app:
	go run cmd/main.go