# System
OS := $(shell uname)

# Docker services
DOCKER_COMPOSE_PROJECT_NAME=exercise_manager_services_local
LOCAL_DOCKER_COMPOSE_FILE=./docker/local/docker-compose.yml

ifeq ($(OS), Linux)
    DOCKER_COMPOSE_COMMAND=docker compose
else
    DOCKER_COMPOSE_COMMAND=docker-compose
endif

up_local_services:
	$(DOCKER_COMPOSE_COMMAND) -p $(DOCKER_COMPOSE_PROJECT_NAME) -f $(LOCAL_DOCKER_COMPOSE_FILE) up -d

down_local_services:
	$(DOCKER_COMPOSE_COMMAND) -p $(DOCKER_COMPOSE_PROJECT_NAME) -f $(LOCAL_DOCKER_COMPOSE_FILE) down

restart_local_services: down_local_services up_local_services

local_services_logs:
	$(DOCKER_COMPOSE_COMMAND) -p $(DOCKER_COMPOSE_PROJECT_NAME) -f $(LOCAL_DOCKER_COMPOSE_FILE) logs

# Migrations(Goose)
MIGRATIONS_DIR=migrations
POSTGRES_DSN=postgresql://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable

apply_migrations:
	goose -dir $(MIGRATIONS_DIR)  postgres "$(POSTGRES_DSN)" up

add_migration:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql
