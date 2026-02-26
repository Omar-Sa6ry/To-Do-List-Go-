# To-Do List API (Go)

A fully functional RESTful API for managing To-Do lists, built using Go (Golang), the Gin framework, PostgreSQL, and JWT authentication. This project serves as a comprehensive implementation of backend best practices and design patterns in Go built after completing Maximilian's Go course.

## Features

- **User Authentication:** Secure user registration and login using bcrypt for password hashing and JSON Web Tokens (JWT) for authentication.
- **RESTful API:** Standardized endpoints for CRUD operations on tasks.
- **Relational Database:** Uses PostgreSQL with `database/sql` for robust data storage and integrity.
- **Middleware:** Custom generic authentication middleware to protect secure routes.
- **Design Patterns:** Implements MVC (Model-View-Controller) architecture, Active Record pattern, and structured project layers.
- **Environment Configuration:** Secure variables management using `.env`.

## Tech Stack

- **Language:** Go (Golang)
- **Framework:** [Gin-Gonic](https://github.com/gin-gonic/gin) (Fast HTTP web framework)
- **Database:** PostgreSQL
- **Database Driver:** [lib/pq](https://github.com/lib/pq)
- **Security:** bcrypt (Password Hashing), [golang-jwt](https://github.com/golang-jwt/jwt) (Authentication Token)
- **Environment Management:** [godotenv](https://github.com/joho/godotenv)

## Project Structure

```text
To Do List (GO)/
├── .env                  # Environment variables (Database credentials, JWT Secret) - NOT commited
├── main.go               # Entry point of the application
├── db/
│   └── db.go             # Database connection and schema initialization
├── models/
│   ├── user.go           # User structure and database operations (Save, Validate)
│   └── todo.go           # Todo structure and CRUD operations
├── middlewares/
│   └── auth.go           # JWT Authentication middleware
├── routes/
│   ├── routes.go         # API routing configuration
│   ├── users.go          # Handlers for Signup and Login
│   └── todos.go          # Handlers for Todo logic
└── utils/
    ├── hash.go           # Password hashing utilities
    └── jwt.go            # JWT generation and verification
```

## Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) installed on your machine.
- PostgreSQL installed and running locally or remotely.

### Installation

1. **Clone the repository:**
   ```bash
   git clone git@github.com:Omar-Sa6ry/To-Do-List-Go-.git
   cd "To Do List (GO)"
   ```

2. **Create Database:**
   Open your PostgreSQL tool and create a new database named `todolist_db` (or any name you prefer).

3. **Configure Environment Variables:**
   Create a `.env` file in the root directory and add the following details:
   ```env
   DB_DRIVER=postgres
   DB_USER=postgres
   DB_PASSWORD=your_postgres_password
   DB_NAME=todolist_db
   DB_HOST=localhost
   DB_PORT=5432
   JWT_SECRET=supersecretkey
   PORT=8080
   ```

4. **Install Dependencies:**
   ```bash
   go mod tidy
   ```

5. **Run the Application:**
   ```bash
   go run .
   ```
   *The server will start on port 8080 (or the port defined in your `.env` file), and database tables will be created automatically.*

## API Documentation (Postman)

A detailed guide on how to test the API endpoints using Postman is available in the `postman_guide.md` file included in this repository.

### Public Endpoints
- `POST /signup` : Create a new user.
- `POST /login` : Authenticate user & get JWT Token.

### Protected Endpoints (Requires Bearer Token)
- `GET /todos` : Retrieve all tasks for the logged-in user.
- `POST /todos` : Create a new task.
- `GET /todos/:id` : Retrieve a specific task.
- `PUT /todos/:id` : Update a specific task.
- `DELETE /todos/:id` : Delete a specific task.

## Author

Omar Sabry
