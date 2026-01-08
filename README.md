# ToDo-APP Auth

This is a simple authentication microservice built with **Go (Golang)**. It allows users to **register** and **log in**, using **JWT** tokens for session management. Perfect as a starting point for projects that need basic authentication.

---

## Technologies Used

- [Golang](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin) - HTTP Framework
- [MySQL](https://www.mysql.com/)
- [GORM](https://gorm.io/index.html) - ORM for Go
- [JWT](https://github.com/golang-jwt/jwt) - Authentication Tokens
- [Docker](https://www.docker.com/)

---

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/dalthonmh/todoapp-auth.git
cd todoapp-auth
```

### 2. Set up the .env file

Create a `.env` file in the project root with the following content:

```bash
DB_DSN=root:password@tcp(mysql:3306)/authdb?charset=utf8mb4 parseTime=True&loc=Local
JWT_SECRET=supersecret
```

### 3. Run the project with Docker

```bash
docker compose up -d --build
```

This will start:

- The Go backend at http://localhost:8080
- A MySQL instance on port 3306

## Endpoints

### POST /register

Register a new user.
Request body:

```json
{
  "username": "user1",
  "password": "123456"
}
```

Response:

```json
{
  "message": "User registered successfully"
}
```

### POST /login

Authenticate a user and return a JWT token.

Request body:

```json
{
  "username": "user1",
  "password": "123456"
}
```

Response:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI..."
}
```
