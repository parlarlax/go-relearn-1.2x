package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, 200, map[string]string{"message": "hello"})
	})

	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		writeJSON(w, 200, map[string]string{"user_id": id})
	})

	mux.HandleFunc("POST /echo", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		w.Write(body)
	})

	mux.HandleFunc("GET /headers", func(w http.ResponseWriter, r *http.Request) {
		headers := map[string]string{}
		for k, v := range r.Header {
			headers[k] = v[0]
		}
		writeJSON(w, 200, headers)
	})

	mux.HandleFunc("GET /query", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		writeJSON(w, 200, map[string]any{"query": q})
	})

	go func() {
		time.Sleep(100 * time.Millisecond)
		demoClient()
	}()

	slog.Info("server on :9999")
	http.ListenAndServe(":9999", mux)
}

func demoClient() {
	fmt.Println("\n=== Client GET ===")
	resp, _ := http.Get("http://localhost:9999/")
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(body))

	fmt.Println("\n=== Client GET with path param ===")
	resp, _ = http.Get("http://localhost:9999/users/42")
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(body))

	fmt.Println("\n=== Client GET with query ===")
	resp, _ = http.Get("http://localhost:9999/query?name=alice&age=30")
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(body))

	fmt.Println("\n=== Client GET headers ===")
	req, _ := http.NewRequest("GET", "http://localhost:9999/headers", nil)
	req.Header.Set("X-Custom", "my-value")
	resp, _ = http.DefaultClient.Do(req)
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(body))

	fmt.Println("\n=== Client POST ===")
	resp, _ = http.Post("http://localhost:9999/echo", "application/json", nil)
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(body))
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
