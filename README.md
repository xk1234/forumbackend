# CVWO Project Backend

This backend was written in Go using the following frameworks:
- GORM as a Postgres wrapper
- Gin web framework to create the REST API
- golang-jwt for authentication

ChatGPT was used to generate API documentation, setup instructions, for debugging, and helped to write some of the authentication code. Postman was used for API testing

# Go/Gin Project with PostgreSQL Setup Instructions

## Prerequisites
- Install [Go](https://golang.org/dl/) (version 1.11 or later for module support).
- Install [PostgreSQL](https://www.postgresql.org/download/).

## Step-by-Step Setup

### 1. Clone Project
```bash
git clone https://github.com/xk1234/cvwobackend.git
```

### 2. Create a Postgres Database
```bash
createdb DB_NAME_HERE
```

### 3. Setup an .env.local file with the correct database credentials
```
DB_HOST=?
DB_USER=?
DB_PASSWORD=?
DB_NAME=?
DB_PORT=?
TOKEN_TTL="2000"
JWT_PRIVATE_KEY="jwt-secret-key"
```


### 4. Run the command below in the project folder
```bash
go run main.go
```

Thats it! The API should now be setup

Test the API with Curl:
```bash
curl -i -H "Content-Type: application/json" \
    -X POST \
    -d '{"username":"nusstudent"}' \
    http://localhost:8000/auth/login
```

# API Documentation



