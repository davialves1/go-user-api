# User Management API

## Overview
This project is a user management REST API built with [Go](https://golang.org/) using the [Gin](https://gin-gonic.com/) framework. It allows the creation, retrieval, searching, and management of user data stored in a database.

---

## Features
- **Create a User**: Add a new user to the database.
- **Get All Users**: Retrieve all users stored in the database.
- **Get a Specific User**: Retrieve the first user from the database.
- **Search Users**: Search for users by name or email.

---

## Project Structure
```
.
├── README.md
├── cmd
│   └── main.go                # Application entry point
├── config
│   └── database.go            # Database connection setup
├── controllers
│   ├── create-user.go         # Controller for user creation
│   └── get-users.go           # Controllers for user retrieval
├── go.mod                     # Go module definition
├── go.sum                     # Dependency checksums
├── models
│   └── user.go                # User model and related methods
└── tests
    └── integration-tests.go   # Integration tests
```

---

## API Endpoints

### Base URL
```
http://localhost:8080
```

### Endpoints

#### **1. Create a User**
- **Endpoint**: `/technical-user`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "email": "user@example.com",
    "name": "John Doe"
  }
  ```
- **Response**:
  ```json
  {
    "id": "<uuid>",
    "email": "user@example.com",
    "name": "John Doe"
  }
  ```

#### **2. Get the First User**
- **Endpoint**: `/technical-user`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "data": {
      "id": "<uuid>",
      "email": "user@example.com",
      "name": "John Doe"
    }
  }
  ```

#### **3. Get All Users**
- **Endpoint**: `/technical-user/all`
- **Method**: `GET`
- **Response**:
  ```json
  [
    {
      "id": "<uuid>",
      "email": "user1@example.com",
      "name": "John Doe"
    },
    {
      "id": "<uuid>",
      "email": "user2@example.com",
      "name": "Jane Doe"
    }
  ]
  ```

#### **4. Search for a User**
- **Endpoint**: `/technical-user/search`
- **Method**: `GET`
- **Query Parameter**:
    - `query`: Search term for email or name
- **Example**:
  ```
http://localhost:8080/technical-user/search?query=john
  ```
- **Response**:
  ```json
  {
    "id": "<uuid>",
    "email": "john.doe@example.com",
    "name": "John Doe"
  }
  ```

---

## Setup and Installation

### Prerequisites
- Go 1.19+
- PostgreSQL

### Steps
1. Clone the repository:
   ```sh
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Set up your environment variables for database configuration in `config/database.go`.

3. Run the application:
   ```sh
   go run cmd/main.go
   ```

4. The server will start on `http://localhost:8080`.

---

## Testing
Integration tests are located in the `tests` folder. To run the tests:
```sh
go test ./tests/...
```

---

## Future Improvements
- Add authentication and authorization.
- Implement pagination for large datasets.
- Improve error handling and validation.
- Add support for environment-based configuration.

---

## License
This project is licensed under the MIT License.

