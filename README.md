# TODO List REST API in Go

This project implements a REST API for creating and managing TODO lists using Go. The application follows the principles of Clean Architecture and utilizes several modern technologies to ensure scalability, maintainability, and ease of use.

## Concepts Covered

- **Web Application Development in Go**: Implementing a RESTful API for creating TODO lists.
- **CRUD Functionality**: Developing Create, Read, Update, and Delete operations for managing TODO lists.
- **Clean Architecture Approach**: Structuring the application following the principles of clean architecture, with well-defined layers.
- **Dependency Injection**: Using dependency injection techniques for managing dependencies within the application.
- **PostgresSQL Database Integration**: Working with a PostgresSQL database to store and retrieve TODO lists.
- **Docker Setup**: Running the application within Docker containers to ensure consistency across different environments.
- **Migration File Generation**: Generating migration files for database schema changes.
- **Application Configuration**: Managing application settings using the `spf13/viper` library and environment variables.
- **Registration and Authentication**: Implementing user registration and authentication, with support for JWT (JSON Web Tokens).
- **Middleware**: Developing middleware for handling authentication, authorization, and other common tasks.
- **SQL Queries**: Writing SQL queries to interact with the database.

## Requirements

- Golang
- Docker
- PostgresSQL

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repository/todo-list-api.git
   cd todo-list-api
