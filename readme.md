# Go CRUD RESTful API

This project is a simple RESTful CRUD API for managing posts, implemented in Go.

## Features

- Built with the Gin framework for efficient HTTP handling.
- Uses GORM as the Object-Relational Mapping library.
- Backed by a PostgreSQL database.
- Comes with a Postman JSON collection for easy API testing.

## Getting Started

### Prerequisites

Ensure you have Go and PostgreSQL installed.

### Database Migration

To set up and migrate the database, run the following command:

```bash
go run migrate/migrate.go
```

## API Testing with Postman

Import the provided Postman collection from the repository to test the API endpoints.

## Auto Compilation

The repository is equipped with CompileDaemon for auto compilation during development. This way, you don't have to run go run main.go every time you make changes.

To initiate auto compilation, use the following command:

```bash
CompileDaemon -command="./go-crud"
```
