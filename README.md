# TaskForge

TaskForge is a REST API built in Go to learn backend development from the ground up using only the Go standard library.

The goal of this project is not just to build an application, but to understand how backend systems work internally - covering HTTP, JSON, routing, request handling, error handling, databases, authentication, testing, and deployment.

---

## Current Features

* HTTP server using `net/http`
* Route handling with `http.ServeMux`
* Health check endpoint
* JSON request and response handling
* Basic error handling
* Method validation

---

## Tech Stack

* Go
* Standard Library (`net/http`, `encoding/json`, etc.)

No third-party frameworks or routers are used.

---

## Project Structure

```text
.
├── go.mod
├── main.go
└── README.md
```

---

## Getting Started

### Prerequisites

* Go 1.24+ (or your installed version)

### Clone the Repository

```bash
git clone https://github.com/jevitapearl/TaskForge.git
cd taskforge
```

### Run the Server

```bash
go run .
```

The server starts on:

```text
http://localhost:8080
```

---

## API Endpoints

### Home

**GET /**

Returns a welcome message.

#### Response

```json
{
  "status": "OK",
  "message": "Welcome to TaskForge"
}
```

---

### Health Check

**GET /health**

Returns the health status of the server.

#### Response

```json
{
  "status": "OK",
  "message": "Health check"
}
```

---

### Echo

**POST /echo**

Returns the message sent in the request body.

#### Request

```json
{
  "message": "Hello, TaskForge!"
}
```

#### Response

```json
{
  "status": "OK",
  "message": "Hello, TaskForge!"
}
```

---

## Example cURL Commands

### Home

```bash
curl http://localhost:8080/
```

### Health

```bash
curl http://localhost:8080/health
```

### Echo

```bash
curl -X POST \
  http://localhost:8080/echo \
  -H "Content-Type: application/json" \
  -d '{"message":"Hello from cURL"}'
```

---

## Goals

By the end of this project, the API will include:

* Authentication
* Task management
* Persistent storage
* Clean architecture
* Unit tests
* Docker support
* Production-ready deployment

---

## License

This project is intended for educational purposes.
