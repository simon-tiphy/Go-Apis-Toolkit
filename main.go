package main

import (
    "encoding/json"
    "log"
    "net/http"
)

// HelloResponse is the JSON shape returned by /hello
type HelloResponse struct {
    Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    resp := HelloResponse{Message: "Hello, World!"}
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        http.Error(w, "failed to encode response", http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    addr := ":8080"
    log.Printf("Server running at http://localhost%s/hello", addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}
