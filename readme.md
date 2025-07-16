# Short Link Project

A simple URL shortener service built with Go, Fiber, and PostgreSQL.

## Features
- Create short links from long URLs
- Redirect to original URLs using short codes
- Get all stored links

## How to Run
1. Clone this repository
2. Copy `.env.example` to `.env` and set your environment variables (DB, BASE_URL, etc)
3. Run database migration (if needed)
4. Build and run the app:
   ```bash
   make run
   # atau
   go run cmd/main.go
   ```

## API Endpoints
- `POST   /v1/shorten`         : Create a new short link
- `GET    /v1/shorten/:id`     : Get detail of a short link by ID
- `GET    /v1/shorten/links`   : Get all short links
- `GET    /v1/:shortCode`      : Redirect to the original URL

## Example Request
```json
POST /v1/shorten
{
  "original_url": "https://example.com"
}
```

## Example Response
```json
{
  "code": 201,
  "data": {
    "id": 10,
    "short_code": "249251",
    "new_link": "localhost:3000/249251",
    "original_url": "https://example.com",
    "created_at": "2025-07-16T06:17:26.380409Z",
    "updated_at": "2025-07-16T06:17:26.380409Z"
  },
  "message": "Short link created successfully",
  "status": 201
}
```

## Tech Stack
- Go
- Fiber
- PostgreSQL

---

> Developed on WSL Fedora. Example IP: `172.18.125.255`