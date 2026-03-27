# URL Shortener Server

A simple URL shortener backend built with **Go**, **Chi router**, and designed to work with a frontend (e.g., Next.js).  
Supports user authentication and link management.

## Features

- User authentication (`/auth/register`, `/auth/login`, `/auth/logout`, `/me`)
- Create short URLs (`POST /api/links`)
- List user links (`GET /api/links`)
- Redirect short URLs (`GET /{slug}`)
- Modular Go project structure

## Project Structure

```
short-url-server/
├── cmd/
│   └── server/
│       └── main.go         # entry point
├── internal/
│   ├── auth/
│   │   ├── handler.go       # auth endpoints
│   │   ├── service.go       # business logic
│   │   ├── repository.go    # DB queries
│   │   └── model.go         # user model
│   ├── links/
│   │   └── handler.go       # create/list links
│   ├── redirect/
│   │   ├── handler.go       # handle redirect
│   │   └── service.go       # redirect logic
│   ├── middleware/
│   │   └── auth.go          # JWT authentication middleware
│   └── store/
│       └── user_store.go    # user DB access
├── go.mod
├── go.sum
└── README.md
````

## Requirements

- Go 1.22+
- PostgreSQL (or any DB of your choice)
- [Chi router](https://github.com/go-chi/chi)

## Setup

1. Clone the repo:

```bash
git clone git@gitlab.com:michelletan1/short-url.git
cd short-url-server
````

2. Install dependencies:

```bash
go mod tidy
```

3. Run the server:

```bash
cd cmd/server
go run main.go
```

Server runs on `http://localhost:8080` by default.

## Endpoints

### Auth

| Method | Path           | Description                |
| ------ | -------------- | -------------------------- |
| POST   | /auth/register | Register a new user        |
| POST   | /auth/login    | Login with email/password  |
| POST   | /auth/logout   | Logout user                |
| GET    | /me            | Get current logged-in user |

### Links

| Method | Path       | Description            |
| ------ | ---------- | ---------------------- |
| POST   | /api/links | Create a new short URL |
| GET    | /api/links | List links for a user  |

### Redirect

| Method | Path    | Description     |
| ------ | ------- | --------------- |
| GET    | /{slug} | Redirect to URL |
