# Go User API

A RESTful API built with Go for user management, providing functionalities for user registration, authentication, and profile management.

## Features

- **User Registration and Authentication:** Secure user sign-up and login.
- **CRUD Operations:** Create, read, update, and delete user profiles.
- **JWT-Based Authentication:** JSON Web Token (JWT) for secure session handling.
- **Docker Support:** Containerized deployment for easy scalability.
- **PostgreSQL Database:** Reliable data storage and retrieval.

## Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.18 or higher recommended)
- [Docker](https://www.docker.com/get-started)
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation

### Clone the Repository

```bash
git clone https://github.com/davialves1/go-user-api.git
cd go-user-api
```

### Set Up Environment Variables

Create a `.env` file based on `.env.example` and update it with your configuration:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
JWT_SECRET=yoursecretkey
```

### Install Dependencies

```bash
go mod download
```

## Project Structure

```
├── cmd
│   └── main.go
├── config
│   └── config.go
├── controllers
│   └── user_controller.go
├── models
│   └── user.go
├── routes
│   └── routes.go
├── services
│   └── user_service.go
├── .env.example
├── docker-compose.yml
├── go.mod
├── go.sum
├── README.md
```

## Running the Application

### Using Docker

1. Build and run the containers:

   ```bash
   docker-compose up --build -d
   ```

2. Access the API at `http://localhost:8000`

### Running Locally

1. Start the application:

   ```bash
   go run main.go
   ```

2. Access the API at `http://localhost:8000`

## API Endpoints

The following endpoints are available:

- **User Registration:** `POST /users/register`
- **User Login:** `POST /users/login`
- **Get User Profile:** `GET /users/{id}`
- **Update User Profile:** `PUT /users/{id}`
- **Delete User:** `DELETE /users/{id}`

Refer to the [API Documentation](https://documenter.getpostman.com/view/6929584/2s8YzTTh9w) for detailed usage.

## Running Tests

To run unit tests:

```bash
go test ./... -v
```

## License

This project is licensed under the MIT License.

---

_Ensure you do not commit sensitive data such as credentials or JWT secrets to the repository._
