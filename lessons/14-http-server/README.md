# บทที่ 14: HTTP Server — net/http router

> 🆕 **Go 1.22+** — `net/http` มี router ใหม่ รองรับ method matching + path parameters โดยไม่ต้องใช้ chi/gin

## Java Spring Boot vs Go net/http

```java
// Spring Boot
@RestController
@RequestMapping("/api")
public class UserController {

    @GetMapping("/users/{id}")
    public ResponseEntity<User> getUser(@PathVariable int id) {
        return ResponseEntity.ok(userService.findById(id));
    }

    @PostMapping("/users")
    public ResponseEntity<User> createUser(@RequestBody User user) {
        return ResponseEntity.ok(userService.save(user));
    }
}
```

```go
// Go 1.22+
mux := http.NewServeMux()

mux.HandleFunc("GET /api/users/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    fmt.Fprintf(w, "User: %s", id)
})

mux.HandleFunc("POST /api/users", func(w http.ResponseWriter, r *http.Request) {
    // handle create
})

http.ListenAndServe(":8080", mux)
```

## สิ่งที่ใหม่ใน Go 1.22+ ServeMux

| ฟีเจอร์ | ก่อน 1.22 | 1.22+ |
|---|---|---|
| Method matching | ต้องเช็ค `r.Method` เอง | `"GET /path"` |
| Path parameters | ต้องใช้ library ภายนอก | `{id}` + `r.PathValue("id")` |
| Wildcard | ไม่มี | `{path...}` (catch-all) |
| Trailing slash | สับสนง่าย | มีกฎชัดเจนขึ้น |

## Handler Pattern

```go
// Simple handler
mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello!")
})

// Handler with path parameter
mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    json.NewEncoder(w).Encode(map[string]string{"id": id})
})

// Method-specific routing
mux.HandleFunc("GET /items", listItems)     // GET only
mux.HandleFunc("POST /items", createItem)   // POST only
mux.HandleFunc("DELETE /items/{id}", deleteItem) // DELETE only
```

## Middleware Pattern

```java
// Spring Boot — @Component + Filter interface
@Component
public class AuthFilter implements Filter {
    public void doFilter(request, response, chain) {
        // before
        chain.doFilter(request, response);
        // after
    }
}
```

```go
// Go — function wrapper
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        slog.Info("request", "method", r.Method, "path", r.URL.Path,
            "duration", time.Since(start))
    })
}

// ใช้
handler := LoggingMiddleware(mux)
http.ListenAndServe(":8080", handler)
```

## JSON Response

```go
func jsonHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]any{
        "status": "ok",
        "data":   []string{"alice", "bob"},
    })
}
```

## ไฟล์ในบทนี้

- `main.go` — CRUD API + middleware + JSON response (รันแล้ว curl ทดสอบได้)
