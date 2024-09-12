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
* Add migration: `make add_migration name=<migration_name>`
  * It will create an empty migration file in `internal/database/migrations` directory. 
* Goose apply migration command for local development: `make apply_migrations`

# Development
1. Download dependencies: `go mod download`
2. Start docker services: `make up_local_services`
3. Install goose for migrations: `go install github.com/pressly/goose/v3/cmd/goose@latest`
