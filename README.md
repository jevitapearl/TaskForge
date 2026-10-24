# TaskForge

TaskForge is a REST API built in Go to learn backend development from the ground up, using the Go standard library for the HTTP layer and a small set of well-known libraries for the pieces the standard library does not cover.

The goal of this project is not just to build an application, but to understand how backend systems work internally — covering HTTP, JSON, routing, request handling, error handling, databases, authentication, authorization, middleware, migrations, and testing.

---

## Current Features

* HTTP server using `net/http` and the Go 1.22+ enhanced `http.ServeMux` (method + path wildcards)
* Layered architecture: `handler → service → repository`
* Health check and echo endpoints
* JSON request/response handling with consistent error payloads
* Method validation and 404 handling
* PostgreSQL persistence with SQL migrations run automatically at startup
* JWT authentication with short-lived access tokens and rotating refresh tokens stored in `HttpOnly` cookies
* bcrypt password hashing
* Per-user task ownership — every task query is scoped to the authenticated user
* Middleware: request logging, panic recovery, and authentication
* Task CRUD with pagination and filtering query parameters

---

## Tech Stack

| Concern | Choice |
| --- | --- |
| Language | Go 1.26 |
| HTTP | `net/http` (standard library) |
| Routing | `http.ServeMux` with method patterns (standard library) |
| Database | PostgreSQL via `database/sql` + `lib/pq` |
| Migrations | `golang-migrate/migrate/v4` (file source) |
| Auth | `golang-jwt/jwt/v5` (HS256) |
| Password hashing | `golang.org/x/crypto/bcrypt` |
| IDs | `google/uuid` + PostgreSQL `pgcrypto` |
| Config | `joho/godotenv` |

No web framework or third-party router is used — routing, middleware, and the request lifecycle are all hand-rolled on top of the standard library.

---

## Project Structure

```text
.
├── cmd/
│   └── server/
│       └── main.go          # entrypoint: config → migrations → db → router → server
├── internal/
│   ├── auth/                # JWT generation/validation, bcrypt hashing
│   ├── config/              # .env loading
│   ├── contextx/            # request-scoped user context
│   ├── database/            # postgres connection setup
│   ├── handler/             # HTTP layer (decode, validate, respond)
│   ├── middleware/          # logging, auth, panic recovery
│   ├── migrations/          # .up.sql / .down.sql migration files
│   ├── models/              # domain types
│   ├── repository/          # data access (postgres + in-memory)
│   ├── router/              # route table
│   └── service/             # business logic
├── .env                     # local config (git-ignored)
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

---

## Getting Started

### Prerequisites

* Go 1.26+ (check with `go version`)
* PostgreSQL 12+ running locally
* `psql` client (optional, for creating the database)

### 1. Clone the Repository

```bash
git clone https://github.com/jevitapearl/TaskForge.git
cd TaskForge
```

### 2. Create the Database

```bash
createuser taskforge
createdb -O taskforge taskforge
```

If you use a different role or database name, just make sure the values in `.env` match your local setup.

### 3. Configure the Environment

Create a `.env` file in the project root:

```bash
DBUSER=taskforge
DBPASSWORD=your_password
HOST=localhost
DBPORT=5432
DBNAME=taskforge
JWT_SECRET=change_me_to_a_long_random_string
```

| Variable | Required | Description |
| --- | --- | --- |
| `DBUSER` | yes | PostgreSQL user |
| `DBPASSWORD` | yes | PostgreSQL password |
| `HOST` | yes | PostgreSQL host |
| `DBPORT` | yes | PostgreSQL port (parsed as an integer) |
| `DBNAME` | yes | Database name |
| `JWT_SECRET` | yes | HMAC secret used to sign access and refresh tokens |
| `PORT` | no | Not read by the server — the HTTP port is fixed at `8080` in `cmd/server/main.go` |

### 4. Run the Server

```bash
go mod tidy
go run ./cmd/server
```

On startup the server applies any pending migrations from `internal/migrations`, opens the database connection, and starts listening on:

```text
http://localhost:8080
```

Verify it is up:

```bash
curl http://localhost:8080/health
```

---

## Database & Migrations

Migrations live in `internal/migrations` and are applied automatically on boot via `golang-migrate` (`file://internal/migrations`). To apply them manually:

```bash
migrate -path internal/migrations -database "postgres://taskforge:your_password@localhost:5432/taskforge?sslmode=disable" up
```

| Version | Description |
| --- | --- |
| `000001_create_users` | `users` table — id, unique email, password hash, role, timestamps |
| `000002_create_tasks` | `tasks` table — id, title, status, owner (`user_id` FK → `users.id`) |
| `000003_create_refresh_tokens` | `refresh_tokens` table — token, user, expiry (cascading delete on user removal) |
| `000004_update_tasks` | adds `created_at` / `updated_at` to `tasks` |

---

## API Endpoints

### Authentication model

* `POST /login` sets two `HttpOnly` cookies: `access_token` (15 minutes) and `refresh_token` (7 days).
* All `/tasks` endpoints require a valid `access_token` cookie; otherwise the server responds `401`.
* `POST /refresh` issues a new access token and rotates the refresh token (the old one is deleted from the database in the same transaction).
* `POST /logout` revokes the stored refresh token and clears both cookies.

### Public endpoints

#### Home

**GET /**

```json
{
  "status": 200,
  "message": "Welcome to TaskForge"
}
```

#### Health Check

**GET /health**

```json
{
  "status": 200,
  "message": "Health check"
}
```

#### Echo

**POST /echo**

Request:

```json
{
  "message": "Hello, TaskForge!"
}
```

Response:

```json
{
  "status": 202,
  "message": "Hello, TaskForge!"
}
```

#### Register

**POST /register**

Request:

```json
{
  "email": "dev@example.com",
  "password": "secret123"
}
```

Response (`201 Created`):

```json
{
  "message": "registered"
}
```

The email must be unique — a duplicate returns `400` with `{"error": "..."}`.

#### Login

**POST /login**

Request:

```json
{
  "email": "dev@example.com",
  "password": "secret123"
}
```

Response (`200 OK`, plus `Set-Cookie` headers):

```json
{
  "message": "logged in"
}
```

Invalid credentials return `401` with `{"error": "invalid credentials"}`.

#### Refresh

**POST /refresh**

Uses the `refresh_token` cookie and rotates it.

Response (`200 OK`):

```json
{
  "message": "token refreshed"
}
```

A missing or expired token returns `401`.

#### Logout

**POST /logout**

Response (`200 OK`):

```json
{
  "message": "logged out"
}
```

---

### Protected endpoints

All endpoints below require the `access_token` cookie obtained from `POST /login`.

#### List Tasks

**GET /tasks**

Query parameters:

| Parameter | Default | Description |
| --- | --- | --- |
| `limit` | `10` | Page size; must be a positive integer |
| `page` | `0` | Zero-based page number; must be `>= 0` |
| `completed` | — | Boolean filter on the task status (`true` / `false`) |
| `sort` | `task_id` | One of `created_at`, `title`, `modified` |

Response (`200 OK`):

```json
[
  {
    "task_id": "5b3b6b52-0d3b-4d0f-9e2a-3f6f2a1c9d84",
    "title": "Write the README",
    "status": false
  }
]
```

Only tasks owned by the authenticated user are returned.

#### Create Task

**POST /tasks**

Request:

```json
{
  "title": "Write the README",
  "status": false
}
```

Response (`201 Created`):

```json
{
  "status": 201,
  "message": "Created"
}
```

Error responses:

| Status | Body | Cause |
| --- | --- | --- |
| `400` | `{"error": "title cannot be empty"}` | Blank title |
| `400` | `{"error": "title too long"}` | Title longer than 100 characters |
| `409` | `{"error": "duplicate title"}` | Title already exists for this user |

#### Get Task

**GET /tasks/{id}**

Response (`200 OK`):

```json
{
  "task_id": "5b3b6b52-0d3b-4d0f-9e2a-3f6f2a1c9d84",
  "title": "Write the README",
  "status": false
}
```

Not found returns `404` with `{"error": "task not found"}`.

#### Update Task

**PUT /tasks/{id}**

Request:

```json
{
  "title": "Write the README (v2)",
  "status": true
}
```

Response (`200 OK`):

```json
{
  "status": 200,
  "message": "Edited"
}
```

#### Delete Task

**DELETE /tasks/{id}**

Response (`200 OK`):

```json
{
  "status": 200,
  "message": "Deleted"
}
```

Not found returns `404` with `{"error": "task not found"}`.

---

## Example cURL Commands

A full flow with a cookie jar (`cookie.txt` is git-ignored):

```bash
# Register a user
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"email":"dev@example.com","password":"secret123"}'

# Log in and persist the auth cookies
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"dev@example.com","password":"secret123"}' \
  -c cookie.txt

# Create a task
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -b cookie.txt \
  -d '{"title":"Write the README","status":false}'

# List tasks
curl "http://localhost:8080/tasks?limit=10&page=0" -b cookie.txt

# List only completed tasks
curl "http://localhost:8080/tasks?completed=true" -b cookie.txt

# Update a task
curl -X PUT http://localhost:8080/tasks/<task_id> \
  -H "Content-Type: application/json" \
  -b cookie.txt \
  -d '{"title":"Write the README (v2)","status":true}'

# Delete a task
curl -X DELETE http://localhost:8080/tasks/<task_id> -b cookie.txt

# Rotate the tokens
curl -X POST http://localhost:8080/refresh -b cookie.txt -c cookie.txt

# Log out
curl -X POST http://localhost:8080/logout -b cookie.txt
```

Public endpoints:

```bash
curl http://localhost:8080/
curl http://localhost:8080/health

curl -X POST \
  http://localhost:8080/echo \
  -H "Content-Type: application/json" \
  -d '{"message":"Hello from cURL"}'
```

---

## Architecture

```text
HTTP request
   │
   ▼
RecoveryMiddleware          panic → 500 instead of a dropped connection
   │
   ▼
LoggingMiddleware           logs method, path, status, duration, user_id
   │
   ▼
AuthMiddleware (protected)  validates the access_token cookie, injects the user into the context
   │
   ▼
Router / Handler            decode + validate the request, write the response
   │
   ▼
Service                     business rules (title length, duplicate detection, token rotation)
   │
   ▼
Repository                  SQL against PostgreSQL
```

Every layer only talks to the one below it. Handlers never touch the database, and the repository never contains business rules.

### Middleware

| Middleware | Purpose |
| --- | --- |
| `RecoveryMiddleware` | Catches panics and returns `500 Internal Server Error` |
| `LoggingMiddleware` | Logs `user_id`, method, path, status code and duration |
| `AuthMiddleware` | Reads the `access_token` cookie, validates the JWT, and stores the user id/role in the request context |

---

## Known Issues

Work in progress, tracked so the docs stay honest:

* The `sort` query parameter is validated but never reaches the repository — the handler parses `sort` into a local variable that is never assigned, so ordering always falls back to `task_id`.
* The `completed` filter is always applied because `TaskFilter.Completed` is a plain `bool` rather than a nullable one — omitting the parameter therefore filters on `status = false` instead of returning every task.
* `PUT /tasks/{id}` returns `500` for an unknown task id, because the handler does not check for `sql.ErrNoRows` the way the delete and get handlers do.
* `POST /echo` returns HTTP `200` while the body reports `"status": 202`.
* The `binding` struct tags on the register payload are decorative — no validator library is wired up, so only the database `UNIQUE` constraint on email is actually enforced.
* There are no unit or integration tests yet.

---

## Goals

Remaining work planned for this project:

* Request validation with a proper validator
* Unit and integration tests with a test database
* Rate limiting and brute-force protection on the auth endpoints
* Pagination metadata in the list response
* Roles and authorization beyond "owns the resource"
* Docker support (app + database) and CI
* Production-ready deployment (graceful shutdown, structured logging, metrics)

---

## License

This project is intended for educational purposes.
