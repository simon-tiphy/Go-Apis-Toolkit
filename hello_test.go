package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/hello", nil)
    w := httptest.NewRecorder()
    helloHandler(w, req)
    res := w.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Fatalf("expected status 200, got %d", res.StatusCode)
    }

    var h HelloResponse
    if err := json.NewDecoder(res.Body).Decode(&h); err != nil {
        t.Fatalf("failed to decode response: %v", err)
    }

    if h.Message != "Hello, World!" {
        t.Fatalf("unexpected message: %s", h.Message)
    }
}
