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

## Protected Routes
These routes require JWT authentication.

### Create Post
- **Method**: POST
- **URL**: `/api/posts`
- **Description**: Allows authenticated users to add new posts.

### Read User Posts
- **Method**: GET
- **URL**: `/api/posts`
- **Description**: Retrieves posts created by the authenticated user.

### Update Post
- **Method**: PUT
- **URL**: `/api/posts/:id`
- **Description**: Allows authenticated users to update their post. The `id` in the URL represents the post ID.

### Delete Post
- **Method**: DELETE
- **URL**: `/api/posts/:id`
- **Description**: Allows authenticated users to delete their post. The `id` represents the post ID.

### Create Comment
- **Method**: POST
- **URL**: `/api/comments/:postId`
- **Description**: Enables authenticated users to add a comment to a post. The `postId` in the URL represents the post ID.

### Read Comments
- **Method**: GET
- **URL**: `/api/comments/:postId`
- **Description**: Retrieves comments for a specific post. The `postId` is the post ID.

### Update Comment
- **Method**: PUT
- **URL**: `/api/comments/:id`
- **Description**: Allows authenticated users to update their comment. The `id` in the URL is the comment ID.

### Delete Comment
- **Method**: DELETE
- **URL**: `/api/comments/:id`
- **Description**: Allows authenticated users to delete their comment. The `id` is the comment ID.

## Public Routes
These routes are available to everyone.

### Sign Up
- **Method**: POST
- **URL**: `/public/signup`
- **Description**: Allows new users to sign up.

### Login
- **Method**: POST
- **URL**: `/public/login`
- **Description**: Authenticates users and returns a JWT for accessing protected routes.

### Get All Posts
- **Method**: GET
- **URL**: `/public/posts`
- **Description**: Retrieves all posts available publicly.

### Get Comments
- **Method**: GET
- **URL**: `/public/comments/:postId`
- **Description**: Fetches comments for a specific post. The `postId` in the URL is the post ID.

