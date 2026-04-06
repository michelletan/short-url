# URL Shortener Backend

A simple URL shortener backend built with **Go**, **Chi router**, and designed to work with a frontend (e.g., Next.js).  
Supports user authentication and link management.

## Features

- User authentication (`/auth/register`, `/auth/login`, `/auth/logout`, `/me`)
- Create short URLs (`POST /api/links`)
- List user links (`GET /api/links`)
- Redirect short URLs (`GET /{slug}`)
- Modular Go project structure

## Requirements

- Go 1.22+
- PostgreSQL (or any DB of your choice)
- [Chi router](https://github.com/go-chi/chi)

## Setup

1. Clone the repo:

```bash
git clone git@gitlab.com:michelletan1/short-url.git
cd backend
````

2. Install dependencies:
```bash
make install-tools
```
or

```bash
go mod tidy
```

3. **Copy the `.env.example` file** to a new `.env.dev` file:

```bash
cp .env.example .env.dev
```

4. Run the server:

Locally
```bash
make run
```
or
```bash
cd cmd/server
go run main.go
```

As a container
```bash
docker-compose up -d --build
```

Server runs on `http://localhost:8080` by default.

## Endpoints

### Auth

| Method | Path           | Description                |
| ------ | -------------- | -------------------------- |
| POST   | /auth/register | Register a new user        |
| POST   | /auth/login    | Login with email/password  |
| POST   | /auth/logout   | Logout user                |
| GET    | /api/me        | Get current logged-in user |

### Links
Requires authentication
| Method | Path       | Description            |
| ------ | ---------- | ---------------------- |
| POST   | /api/links | Create a new short URL |
| GET    | /api/links | List links for a user  |

### Redirect

| Method | Path    | Description     |
| ------ | ------- | --------------- |
| GET    | /{slug} | Redirect to URL |
