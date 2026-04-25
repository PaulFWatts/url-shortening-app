# Copilot Instructions

## Build, Test, Lint
- Run server: `go run cmd\web-app\main.go`
- Build binary: `go build -o web-app cmd\web-app\main.go`
- Test all packages (when tests exist): `go test ./...`
- Run a single test: `go test ./... -run TestName`

## High-level Architecture
- Entry point is `cmd/web-app/main.go`, which registers two HTTP handlers and starts an `http.ListenAndServe` on port 8080.
- Request handling lives in `internal/controllers` with `ShowIndex` for `GET /` and `Shorten` for `POST /shorten`.
- HTML templates are rendered server-side from `internal/views` using `html/template` and returned as full pages or fragments.
- The index page uses client-side JavaScript to POST to `/shorten` and injects the returned HTML into the page.

## Key Conventions
- Handlers parse templates by relative path (e.g., `internal/views/index.html`), so running from repo root is expected.
- The shorten flow normalizes URLs by adding `https://` when no scheme is provided.
- The `Shorten` handler currently returns the original URL as `ShortURL` in the template data map.
- For utility scripts (for example, `scripts\myenv\gen_seo.py`), resolve project files from `__file__`/absolute paths instead of process cwd-relative paths.
- Never commit API keys or secrets in code; read them from environment variables.
