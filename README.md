# Markdown API Backend

A Go backend API built with **Gin** and **Supabase** that allows users to:

- Register and log in
- Authenticate and authorize with JWT
- Create and manage Markdown documents

---

## Features

- 🔐 User Authentication (JWT)
- ✅ User Authorization
- 📝 Create and Send Markdown documents
- 🚀 Built with Go, Gin, and Supabase (Postgres)

---

## Tech Stack

- **Go** - Backend programming language
- **Gin** - Web framework
- **Supabase** - Postgres database and user storage
- **JWT** - Authentication mechanism

---

## Setup Instructions

### Prerequisites

- Go 1.20 or higher
- A [Supabase](https://supabase.io/) project (set up database and authentication)

### Clone the Repository

```bash
git clone https://github.com/Ik-cyber/markdown-ai-app.git
cd markdown-ai-app
```

### Environment Variables

Create a `.env` file in the root directory with the following keys:

```env
SUPABASE_URL=your-supabase-url
SUPABASE_API_KEY=your-supabase-api-key
JWT_SECRET=your-jwt-secret-key
```

### Install Dependencies

```bash
go mod tidy
```

### Run the Server

```bash
go run main.go
```

The server will start on:

```
http://localhost:7777
```

---

## API Usage

### Auth Routes

#### Register

**POST** `/api/v1/register`

Request Body:

```json
{
  "username": "your-username",
  "password": "your-password"
}
```

#### Login

**POST** `/api/v1/login`

Request Body:

```json
{
  "username": "your-username",
  "password": "your-password"
}
```

Response:

```json
{
  "token": "your-jwt-token"
}
```

---

### Markdown Routes (Protected)

> Include the token in the `Authorization` header:

```http
Authorization: Bearer <JWT_TOKEN>
```

#### Create Markdown

**POST** `/api/v1/markdowns`

Request Body:

```json
{
  "title": "Sample Title",
  "content": "# Hello World\nThis is my first markdown."
}
```

#### Get All Markdowns

**GET** `/api/v1/markdowns`

#### Get Single Markdown

**GET** `/api/v1/markdowns/:id`

#### Update Markdown

**PUT** `/api/v1/markdowns/:id`

Request Body:

```json
{
  "title": "Updated Title",
  "content": "# Updated Markdown Content"
}
```

#### Delete Markdown

**DELETE** `/api/v1/markdowns/:id`

---

## Example Curl Commands

### Register

```bash
curl -X POST http://localhost:7777/api/v1/register \
-H "Content-Type: application/json" \
-d '{"username":"testuser", "password":"testpass"}'
```

### Login

```bash
curl -X POST http://localhost:7777/api/v1/login \
-H "Content-Type: application/json" \
-d '{"username":"testuser", "password":"testpass"}'
```

### Create Markdown

```bash
curl -X POST http://localhost:7777/api/v1/markdowns \
-H "Authorization: Bearer <JWT_TOKEN>" \
-H "Content-Type: application/json" \
-d '{"title":"First Markdown", "content":"# Hello World\nThis is my first markdown."}'
```

---

## Project Structure

```text
.
├── api
│   └── v1
│       ├── auth.go           # Handles registration and login
│       ├── markdown.go       # Handles markdown CRUD operations
│       └── routes.go         # Registers API routes
├── internals
│   └── database
│       └── db.go             # Supabase client setup and database logic
├── middleware
│   └── auth.go               # JWT Authentication middleware
├── go.mod
├── go.sum
├── main.go                   # App entry point
├── .env                      # Environment variables
└── README.md                 # Project documentation
```

---

## How to Use

1. Start the server:

   ```bash
   go run main.go
   ```

2. Register a user via the `/api/v1/register` endpoint.

3. Login via the `/api/v1/login` endpoint to receive a JWT token.

4. Use the token in the `Authorization` header for all Markdown routes.

5. Create, fetch, update, and delete Markdown documents using the provided API endpoints.
