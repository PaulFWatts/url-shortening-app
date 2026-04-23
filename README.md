# Url Shortening App

A simple Go-based web application for shortening URLs. This project is part of the [Mastering Go With GoLand](https://www.bytesizego.com/courses/a508e491-8e7b-4f29-9cbb-a374126cecac/view) course.

## Overview

The application provides a web interface where users can enter a URL and receive a shortened version. 
Currently, it serves as a foundational project demonstrating:
- Basic Go web server using `net/http`.
- HTML template rendering with `html/template`.
- Handling form submissions.

## Requirements

- **Go**: 1.26 or higher (as specified in `go.mod`).

## Setup and Run

### 1. Clone the repository
```bash
git clone https://github.com/PaulFWatts/url-shortening-app.git
cd url-shortening-app
```

### 2. Run the application
You can start the web server directly using the Go tool:
```bash
go run cmd/web-app/main.go
```
The server will start on `http://localhost:8080`.

## Scripts

- `go run cmd/web-app/main.go`: Starts the development server.
- `go build -o web-app cmd/web-app/main.go`: Builds the executable.

## Project Structure

```text
url-shortening-app/
├── cmd/
│   └── web-app/
│       └── main.go         # Application entry point
├── internal/
│   ├── controllers/
│   │   ├── index.go        # Handlers for the index page
│   │   └── shorten.go      # Handlers for URL shortening logic
│   └── views/
│       ├── index.html      # Home page template
│       └── shorten.html    # Results page template
├── go.mod                  # Go module definition
└── README.md               # Project documentation
```

## Environment Variables

Currently, the application does not use any environment variables.
- TODO: Add support for port configuration via `PORT` environment variable.

## Tests

- TODO: Add unit tests for controllers and shortening logic.
- Run tests (once added) using: `go test ./...`

## License

- TODO: Add license information.
