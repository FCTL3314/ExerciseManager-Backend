# ExerciseManager Backend

Backend server that provides an API for workout reminders and tracking exercise completion, including fitness activities like calisthenics and so on. 

Goose base command: `goose -dir db/migrations <command>`
Goose local migration command: `goose -dir internal/database/migrations postgres "postgresql://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable" up`

# Development
1. Download dependencies: `go mod download`
2. Start docker services: `docker-compose -p exercise_manager_local -f .\docker\local\docker-compose.yml up -d`
3. Install goose for migrations: `go install github.com/pressly/goose/v3/cmd/goose@latest`
