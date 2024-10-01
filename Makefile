.PHONY: update

# System
OS := $(shell uname)

# Docker services
DOCKER_COMPOSE_PROJECT_NAME=exercise_manager_backend_services_local
LOCAL_DOCKER_COMPOSE_FILE=./docker/local/docker-compose.yml

up_local_services:
	docker compose -p $(DOCKER_COMPOSE_PROJECT_NAME) -f $(LOCAL_DOCKER_COMPOSE_FILE) up -d

down_local_services:
	docker compose -p $(DOCKER_COMPOSE_PROJECT_NAME) -f $(LOCAL_DOCKER_COMPOSE_FILE) down

restart_local_services: down_local_services up_local_services

local_services_logs:
	docker compose -p $(DOCKER_COMPOSE_PROJECT_NAME) -f $(LOCAL_DOCKER_COMPOSE_FILE) logs

# Migrations(Goose)
MIGRATIONS_DIR=migrations
POSTGRES_DSN_DEFAULT=postgresql://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable

apply_migrations:
	goose -dir $(MIGRATIONS_DIR)  postgres "$(or $(POSTGRES_DSN), $(POSTGRES_DSN_DEFAULT))" up

add_migration:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

# Deployment
DEFAULT_SERVICE_NAME = exercise_manager_server

get_updates:
	git pull
	sudo systemctl restart $(or $(SERVICE_NAME), $(DEFAULT_SERVICE_NAME))
	sudo systemctl --no-pager -l status $(or $(SERVICE_NAME), $(DEFAULT_SERVICE_NAME))

