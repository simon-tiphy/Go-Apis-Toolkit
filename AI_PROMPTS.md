# AI Prompts Journal (sample) — replace with your actual prompts & reflections

1) Prompt:
```
Show me a minimal Go HTTP server that returns JSON 'Hello, World!' on GET /hello. Include code and instructions to run.
```
AI Response Summary:
- Provided a `net/http` example and explained `go run`.
Reflection:
- Helped scaffold handler and JSON encoding. Needed a follow-up about Content-Type header.

2) Prompt:
```
How do I write a unit test for an HTTP handler in Go using net/http/httptest?
```
AI Response Summary:
- Suggested using httptest.NewRecorder and NewRequest, then decode JSON response.
Reflection:
- Test code worked first try; adjusted imports per ai suggestion.

3) Prompt:
```
What are common pitfalls when running a Go server locally (like port in use) and how to resolve them?
```
AI Response Summary:
- Change port, kill process using the port, check firewall.
Reflection:
- Resolved a port conflict by switching to :8081 briefly.

---

