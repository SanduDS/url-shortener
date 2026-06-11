# URL Shortener

A simple URL shortener service written in Go. It exposes a small HTTP API that takes a long URL and returns a shortened one, using an in-memory store and base62-encoded short codes.

## Project Structure

```
.
├── cmd/
│   └── server/         # Application entry point (HTTP server)
├── internal/
│   ├── handler/        # HTTP handlers
│   ├── model/          # Request/response and domain models
│   └── service/        # Business logic (in-memory URL store)
├── pkg/
│   └── base62/         # Base62 encoding for short codes
├── migrations/         # Database migrations (placeholder)
├── docs/               # Documentation (placeholder)
└── go.mod
```

## Requirements

- Go 1.25+

## Getting Started

Clone the repository and run the server:

```bash
git clone https://github.com/SanduDS/url-shortener.git
cd url-shortener
go run ./cmd/server
```

The server starts on port `8080`.

## API

### Shorten a URL

```
POST /api/v1/shorten
```

Request body:

```json
{
  "url": "https://example.com/some/very/long/path"
}
```

Response:

```json
{
  "short_url": "http://localhost:8080/1"
}
```

Example with `curl`:

```bash
curl -X POST http://localhost:8080/api/v1/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com/some/very/long/path"}'
```

### Health Check

```
GET /health
```

Returns `200 OK` when the service is up.

## How It Works

Each shortened URL gets an incrementing counter value that is encoded with [base62](pkg/base62/base62.go) (`0-9A-Za-z`) to produce a compact short code. URLs are stored in an in-memory map, so data is lost on restart.

## Roadmap

- Redirect endpoint to resolve short codes back to original URLs
- Persistent storage (database) with migrations
- Configurable base URL and port
