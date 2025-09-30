# Go Hello API — Beginner Toolkit (Prompt-Powered Capstone)

This is a minimal Go project for the Moringa AI Capstone: a small HTTP API that returns JSON `{"message":"Hello, World!"}`.

## Requirements
- Go 1.20+ installed
- Git (optional)

## Files
- `main.go` — minimal HTTP server with `/hello`
- `hello_test.go` — unit test for the handler
- `go.mod` — module file (replace `github.com/your-username/...` with your path)
- `AI_PROMPTS.md` — sample AI prompts & reflections (edit with your real prompts)
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

## Notes
- Replace module path in `go.mod` with your GitHub path before publishing.
- Use `go fmt` and `go vet` to lint and format code.
