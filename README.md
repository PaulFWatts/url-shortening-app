# URL Shortening App

A small Go web app that accepts a URL and returns a deterministic short code.

## Current behavior

- Serves a web form at `GET /`.
- Accepts shortening requests at `POST /shorten`.
- Normalizes input by prepending `https://` when no scheme is provided.
- Generates short codes using SHA-256 and returns the first 8 hex characters.
- Renders the shortened result as an HTML fragment and injects it into the page with JavaScript.
- Includes a copy-to-clipboard button for `http://localhost:8080/<short-code>`.

## Requirements

- **Go** `1.26.1` (see `go.mod`).

## Run

```bash
go run cmd/web-app/main.go
```

App URL: `http://localhost:8080`

## Build

```bash
go build -o web-app cmd/web-app/main.go
```

## Test

```bash
go test ./...
```

## SEO helper script

The repo includes `scripts\myenv\gen_seo.py` to generate SEO text and append it to `internal\views\index.html`.

- The script now resolves `index.html` from the script location, so it is not tied to your current working directory.
- Do **not** hardcode API keys in source files. Use an environment variable instead.

PowerShell example:

```powershell
$env:OPENAI_API_KEY="your-key"
python scripts\myenv\gen_seo.py
```

## Project structure

```text
url-shortening-app/
├── cmd/
│   └── web-app/
│       └── main.go                 # Registers routes and starts :8080
├── internal/
│   ├── controllers/
│   │   ├── index.go                # GET / handler
│   │   └── shorten.go              # POST /shorten handler
│   └── views/
│       ├── index.html              # Form + client-side fetch logic
│       └── shorten.html            # Result fragment with copy button
├── url/
│   └── url.go                      # Short code generation (SHA-256 -> 8 chars)
├── go.mod
└── README.md
```
