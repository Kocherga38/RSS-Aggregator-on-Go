COMPOSE_FILE = docker-compose.yaml

.PHONY: up down restart logs ps build clean

build:
	go build && rssagg.exe

up:
	docker-compose -f $(COMPOSE_FILE) up -d

down:
	docker-compose -f $(COMPOSE_FILE) down

restart:
	docker-compose -f $(COMPOSE_FILE) down && docker-compose -f $(COMPOSE_FILE) up -d

ps:
	docker-compose -f $(COMPOSE_FILE) ps

clean:
	docker-compose -f $(COMPOSE_FILE) down -v

psql:
	docker exec -it postgres-container psql -U myuser -d mydatabase

