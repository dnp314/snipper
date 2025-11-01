# SNIPPER

A small Go web application that demonstrates user accounts and snippet management using server-rendered HTML templates.

Key files
- [cmd/web/main.go](cmd/web/main.go) — application entry point and server setup
- [cmd/web/routes.go](cmd/web/routes.go) — HTTP route configuration
- [cmd/web/handlers.go](cmd/web/handlers.go) — request handlers
- [cmd/web/middleware.go](cmd/web/middleware.go) — HTTP middleware
- [internal/models/snippets.go](internal/models/snippets.go) — snippet model and DB logic
- [internal/models/users.go](internal/models/users.go) — user model
- [ui/html/base.tmpl](ui/html/base.tmpl) — base HTML template
- [ui/html/pages/home.tmpl](ui/html/pages/home.tmpl) — example page template
- [tls/cert.pem](tls/cert.pem) and [tls/key.pem](tls/key.pem) — local TLS cert/key for HTTPS dev

What it does
- Serves HTML pages using Go templates.
- Provides basic user signup/login flows.
- Lets authenticated users create, view, and list text snippets.
- Includes simple middleware, request context helpers, and unit tests.

Quick start (development)
1. Install dependencies and build:
   ```sh
   go mod download
   go build ./cmd/web
   ```
2. Run the server:
   ```sh
   ./web
   ```
   The server uses the TLS files in [tls/](tls/) for HTTPS in development.

Running tests
- Run all tests:
  ```sh
  go test ./...
  ```
- Specific tests for handlers and middleware live in [cmd/web/handlers_test.go](cmd/web/handlers_test.go) and [cmd/web/middleware_test.go](cmd/web/middleware_test.go).

Development notes
- Templates are under [ui/html/](ui/html/) with partials in [ui/html/partials/](ui/html/partials/).
- Static assets are in [ui/static/](ui/static/) (css, js, images).
- Helper and context utilities are in [cmd/web/helpers.go](cmd/web/helpers.go) and [cmd/web/context.go](cmd/web/context.go).

