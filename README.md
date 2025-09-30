# Go Hello API — Beginner Toolkit

This is a minimal Go project : a small HTTP API that returns JSON `{"message":"Hello, World!"}`.

## Requirements
- Go 1.20+ installed
- Git

## Files
- `main.go` — minimal HTTP server with `/hello`
- `hello_test.go` — unit test for the handler
- `go.mod` — module file github.com/simon-tiphy/Go-Apis-Toolkit
- `AI_PROMPTS.md` — sample AI prompts & reflections
- `.gitignore`

## Run locally
```bash
# from project root
go run .
```

Then open http://localhost:8080/hello or run:
```bash
curl http://localhost:8080/hello
# Expected: {"message":"Hello, World!"}
```

## Run tests
```bash
go test ./...
```

