# CVWO Project Backend

The backend is hosted at https://forumapi-dp4z.onrender.com. To get all posts for example, try:
https://forumapi-dp4z.onrender.com/public/posts

This backend was written in Go using the following frameworks:
- GORM as a Postgres wrapper
- Gin web framework to create the REST API
- golang-jwt for authentication

ChatGPT was used to generate API documentation, setup instructions, for debugging, and helped to write some of the authentication code. Postman was used for API testing

## Prerequisites
- Install [Go](https://golang.org/dl/) (version 1.21.4 or later).
- Install [PostgreSQL](https://www.postgresql.org/download/) and psql command line

## Step-by-Step Setup

### 1. Clone Project
```bash
git clone https://github.com/xk1234/forumbackend.git
cd forumbackend
```

### 2. Create a Postgres Database(enter password if needed)
```bash
psql user=postgres
create database DB_NAME;
\c DB_NAME
\conninfo
\password
```

### 3. Setup an .env.local file with the correct database credentials using the output of conninfo. 
Etc: You are connected to database "testdb" as user "postgres" via socket in "/tmp" at port "5432".
```
DB_HOST="localhost"
DB_USER="postgres"
DB_PASSWORD="YOUR_SET_PASSWORD"
DB_NAME="testdb"
DB_PORT="5432"
TOKEN_TTL="2000"
JWT_PRIVATE_KEY="jwt-secret-key"
```
Alternatively, use the online database
```
DB_HOST="dpg-cmnstoud3nmc739gqgtg-a.oregon-postgres.render.com"
DB_USER="pgres"
DB_PASSWORD="yDuo1opbf1KwIOYu7IcqBeIWKFJZagyT"
DB_NAME="postgres_bqin"
DB_PORT="5432"
TOKEN_TTL="2000"
JWT_PRIVATE_KEY="jwt-secret-key"
```

### 4. Test the connection
```bash
PGPASSWORD="YOUR_SET_PASSWORD" psql -h "localhost" -U "postgres" -d "testdb" -p "5432"
```

### 5. Set sslmode=disable in database/db.go(only for local database)
```go
dbstr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
```


### 6. Run the command below in the project folder
```bash
go run main.go
```

Thats it! The API should now be setup at localhost:4000

Test the API with Curl:
Sign Up
```bash
curl -i -H "Content-Type: application/json" \
    -X POST \
    -d '{"username":"nusstudent"}' \
    http://localhost:4000/public/signup
```

Get Posts
```bash
curl -i -H "Content-Type: application/json" \
    -X GET \
    http://localhost:4000/public/posts
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

