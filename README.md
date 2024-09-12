# ExerciseManager Backend

Backend server that provides an API for workout reminders and tracking exercise completion, including fitness activities like calisthenics and so on. 

# Dependencies
* **GoLang**
* **Docker**
* **Makefile**
  * In Windows can be installed with: `choco install make`

# Notes
* All Docker volumes are stored in the `docker/local/volumes/` folder. If you want to clear your DB or any other data, you can simply delete the folder there.

# Work with migrations
* Goose base command: `goose -dir internal/database/migrations <command>`
* Goose apply migration command for local development: `goose -dir internal/database/migrations postgres "postgresql://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable" up`

# Development
1. Download dependencies: `go mod download`
2. Start docker services: `docker-compose -p exercise_manager_local -f .\docker\local\docker-compose.yml up -d`
3. Install goose for migrations: `go install github.com/pressly/goose/v3/cmd/goose@latest`
