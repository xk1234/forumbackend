# CVWO Project Backend

This backend was written in Go using the following frameworks:
- GORM as a Postgres wrapper
- Gin web framework to create the REST API
- golang-jwt for authentication

ChatGPT was used to generate API documentation, setup instructions, for debugging, and helped to write some of the authentication code. Postman was used for API testing

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

## Protected **Route**s (Require JWT Authentication)
### Add Post

**Method**: POST

**Route**: /api/posts

**Description**: Adds a new post. Requires JWT authentication.
### Update Post

**Method**: PUT

**Route**: /api/posts/:id

**Description**: Updates an existing post by its ID. Requires JWT authentication.
### Delete Post

**Method**: DELETE

**Route**: /api/posts/:id

**Description**: Deletes a post by its ID. Requires JWT authentication.

### Get User Posts

**Method**: GET

**Route**: /api/posts

**Description**: Retrieves all posts created by the authenticated user.

### Add Comment

**Method**: POST

**Route**: /api/comments/:pid

**Description**: Adds a comment to a post specified by the post ID (pid). Requires JWT authentication.

### Update Comment

**Method**: PUT

**Route**: /api/comments/:id

**Description**: Updates a specific comment by its ID. Requires JWT authentication.


### Delete Comment

**Method**: DELETE

**Route**: /api/comments/:id

**Description**: Deletes a specific comment by its ID. Requires JWT authentication.
## Public **Route**s (No Authentication Required)
### Sign Up

**Method**: POST

**Route**: /public/signup

**Description**: Registers a new user.
### Login

**Method**: POST

**Route**: /public/login

**Description**: Authenticates a user and returns a JWT for accessing protected **Route**s.


### Get All Posts

**Method**: GET

**Route**: /public/posts

**Description**: Retrieves all posts.
Get Comments of a Post

**Method**: GET

**Route**: /public/comments/:pid

**Description**: Retrieves all comments for a specific post, identified by the post ID (pid).
### Get Single Post

**Method**: GET

**Route**: /public/post/:id

**Description**: Retrieves details of a single post by its ID.

