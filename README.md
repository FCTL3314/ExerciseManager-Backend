<div align="center">
  <img width="148" height="148" src="https://github.com/user-attachments/assets/b8b8f3ba-d6da-414e-b5f5-339578b498a8"/>
  <h1>Exercise Manager - Backend</h1>
  <p>Backend server for managing workout routines and tracking exercise progress. Supports creating custom workout plans, adding exercises with durations and breaks, and tracking completion of sport activities.</p>

[![Go](https://img.shields.io/badge/Go-1.23.1-45a3ec?style=flat-square)](https://go.dev/)
[![Gin](https://img.shields.io/badge/Gin-1.10.0-458cec?style=flat-square)](https://github.com/gin-gonic/gin)
[![Gorm](https://img.shields.io/badge/Gorm-1.25.12-38B6FF?style=flat-square)](https://github.com/go-gorm/gorm)
[![Viper](https://img.shields.io/badge/Viper-1.19.0-66BC67?style=flat-square)](https://github.com/spf13/viper)
</div>

# 📃 Notes

* All Docker volumes are stored in the `docker/local/volumes/` directory. If you need to reset your database or any other data, simply delete the corresponding folder.

# ⚒️ Development

1. Download dependencies: `go mod download`
2. Install goose for migrations: `go install github.com/pressly/goose/v3/cmd/goose@latest`
3. Start Docker services: `make up_local_services`
