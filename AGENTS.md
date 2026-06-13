# Agent guidance for ~/projects/sample


## Project
A demo Go HTTP server. Single binary. No frameworks.


## Conventions
- Go 1.22+
- Standard library only — no external deps
- All HTTP handlers in `handlers/`
- One file per resource: `users.go`, `posts.go`, etc.
- Test files alongside source: `users_test.go`
- Run with `go run ./cmd/server`
- Build with `go build -o ./bin/server ./cmd/server`


## Don't
- Don't introduce a router library (chi, gorilla/mux, etc.) — stdlib only
- Don't add a database — this demo is in-memory
