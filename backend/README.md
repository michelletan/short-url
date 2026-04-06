# URL Shortener Backend

A simple URL shortener backend built with **Go**, **Chi router**, and designed to work with a frontend (e.g., Next.js).  
Supports user authentication and link management.

## Features

- User authentication (`/auth/register`, `/auth/login`, `/auth/logout`, `/me`)
- Create short URLs (`POST /api/links`)
- List user links (`GET /api/links`)
- Redirect short URLs (`GET /{slug}`)
- Modular Go project structure

## Project Structure

- `cmd/server/`: Main application entry point
- `internal/`: Internal packages (not exported)
  - `config/`: Application configuration
  - `db/`: Database connection and error handling
  - `dtos/`: Data transfer objects for API requests/responses
  - `handlers/`: HTTP request handlers
  - `middleware/`: HTTP middleware (e.g., authentication)
  - `models/`: Data models (User, Link, etc.)
  - `service/`: Business logic services
  - `store/`: Data access layer (repositories)
  - `util/`: Utility functions
  - `validation/`: Input validation logic
- `migrations/`: Database migration scripts
- `bruno/`: API testing collection for Bruno client
- `bin/`: Compiled binaries (e.g., migrate tool)

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

Start the PostgreSQL container:
```bash
make db-up
```

Run database migrations:
```bash
make migrate-up
```

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

## Testing

To test the API endpoints, use [Bruno](https://www.usebruno.com/), an open-source API client.

The Bruno collection is located in the `bruno/` directory. Import the collection into Bruno and run the requests to verify the API functionality.

The collection includes tests for authentication, link creation, listing, and redirection.

## Database

To create a new migration:
```bash
make migrate-create name=your_migration_name
```

To rollback the last migration:
```bash
make migrate-down
```

To stop the database:
```bash
make db-down
```
