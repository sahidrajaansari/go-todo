.PHONY: up down

start:
	docker-compose up --build -d --remove-orphans

stop:
	docker-compose down