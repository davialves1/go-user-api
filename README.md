# Technical User Management API

This project is a simple RESTful API built using the **Go programming language** with the **Gin framework** and **PostgreSQL** as the database. It provides functionality for managing technical users, including creating and retrieving user data.

**This dummy project is part of my learning Go path**

---

## Features

- **Create Technical User**: Add a new technical user to the database.
- **Get Technical User by Email**: Retrieve a user's details using their email address.
- Structured and modular design.
- Uses **GORM** for database interactions.

---

## Prerequisites

To run this project, you need:

- Go (version 1.18 or later)
- PostgreSQL (configured and running)

---

## Setup

### Clone the Repository

```bash
git clone https://github.com/davialves1/go-user-api
cd go-user-api
```

### Configure the Environment

Create a `.env` file in the root directory and provide the following details:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
```

### Install Dependencies

```bash
go mod tidy
```

---

## Running the Application

### Start the Server

Run the following command to start the API server:

```bash
go run main.go
```

The server will start at `http://localhost:8080` by default.

---

## Endpoints

### 1. **Create Technical User**

**Endpoint:** `POST /technical-user`

**Request Body:**

```json
{
  "gid": "string",
  "email": "string",
  "name": "string"
}
```

**Response:**

```json
{
  "data": {
    "id": "uuid",
    "gid": "string",
    "email": "string",
    "name": "string",
    "createdAt": "timestamp"
  }
}
```

### 2. **Get Technical User by Email**

**Endpoint:** `GET /technical-user?email=<email>`

**Response:**

```json
{
  "data": {
    "id": "uuid",
    "gid": "string",
    "email": "string",
    "name": "string",
    "createdAt": "timestamp"
  }
}
```

---

## Project Structure

```plaintext
├── controllers     # API handler functions
├── models          # Database models
├── config          # Database and environment configuration
├── main.go         # Entry point of the application
└── go.mod          # Dependency management
```

---

## Database Schema

### `technical_users` Table

| Column     | Type      | Description           |
|------------|-----------|-----------------------|
| id         | UUID      | Primary Key           |
| gid        | String    | Group Identifier      |
| email      | String    | User Email            |
| name       | String    | User Name             |
| created_at | Timestamp | Record Creation Time  |

---

## Development Notes

- **Testing**:
  Run tests using:

  ```bash
  go test ./...
  ```

- **Updating Dependencies**:
  Use:

  ```bash
  go mod tidy
  go get -u ./...
  ```

---

## License

This project is licensed under the MIT License.

---

## Contributors

- [Davi Alves](https://github.com/davialves1)

---

## Acknowledgements

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)

