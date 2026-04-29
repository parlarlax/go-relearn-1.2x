// gontainer + Layered Architecture example
//
// Structure (horizontal layers, no domain isolation):
//
//	model/       ← data structs (shared across all layers)
//	repository/  ← data access (Spring Boot: @Repository)
//	service/     ← business logic (Spring Boot: @Service)
//	handler/     ← HTTP handlers (Spring Boot: @RestController)
//	main.go      ← gontainer wiring (Spring Boot: @Configuration)
//
// Compare with ddd/ example:
//   - Layered: model/ is shared, no interfaces between layers, simpler but more coupled
//   - DDD:     domain/ owns entities + interfaces, infrastructure implements them, cleaner boundaries
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/NVIDIA/gontainer/v2"

	"github.com/lax/go-relearn/examples/third-party/gontainer/layer/handler"
	"github.com/lax/go-relearn/examples/third-party/gontainer/layer/repository"
	"github.com/lax/go-relearn/examples/third-party/gontainer/layer/service"
)

func main() {
	fmt.Println("=== gontainer + Layered Architecture ===")
	fmt.Println()
	fmt.Println("Layers: handler → service → repository → model")
	fmt.Println("DI:     gontainer auto-injects by concrete type")
	fmt.Println()

	err := gontainer.Run(
		// ── Repository Layer ──────────────────────────────
		gontainer.NewFactory(repository.NewUserRepository),
		gontainer.NewFactory(repository.NewOrderRepository),

		// ── Service Layer ─────────────────────────────────
		gontainer.NewFactory(service.NewUserService),
		gontainer.NewFactory(service.NewOrderService),

		// ── Handler Layer ─────────────────────────────────
		gontainer.NewFactory(handler.NewUserHandler),
		gontainer.NewFactory(handler.NewOrderHandler),

		// ── Server ────────────────────────────────────────
		gontainer.NewFactory(func(
			uh *handler.UserHandler,
			oh *handler.OrderHandler,
		) (*http.Server, func() error) {
			mux := http.NewServeMux()
			uh.RegisterRoutes(mux)
			oh.RegisterRoutes(mux)

			srv := &http.Server{Addr: ":18081", Handler: mux}
			return srv, func() error {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				return srv.Shutdown(ctx)
			}
		}),

		// ── Entrypoint ────────────────────────────────────
		gontainer.NewEntrypoint(func(srv *http.Server) {
			go srv.ListenAndServe()
			waitForServer(":18081")

			base := "http://localhost:18081"
			fmt.Println("--- HTTP Demo Requests ---")
			fmt.Println()

			fmt.Println("POST /users (create Alice)")
			post(base+"/users", `{"name":"Alice","email":"alice@test.com"}`)

			fmt.Println("POST /users (create Bob)")
			post(base+"/users", `{"name":"Bob","email":"bob@test.com"}`)

			fmt.Println("GET /users")
			get(base + "/users")

			fmt.Println("POST /orders (for Alice)")
			post(base+"/orders", `{"user_id":1,"item":"Go Book","qty":3}`)

			fmt.Println("GET /orders")
			get(base + "/orders")

			fmt.Println()
			fmt.Println("--- Cleanup ---")
		}),
	)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done!")
}

func get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("  ERROR: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("  %d %s\n", resp.StatusCode, indentJSON(body))
}

func post(url, payload string) {
	resp, err := http.Post(url, "application/json", strings.NewReader(payload))
	if err != nil {
		fmt.Printf("  ERROR: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("  %d %s\n", resp.StatusCode, indentJSON(body))
}

func indentJSON(raw []byte) string {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return string(raw)
	}
	pretty, err := json.MarshalIndent(v, "  ", "  ")
	if err != nil {
		return string(raw)
	}
	return string(pretty)
}

func waitForServer(addr string) {
	for i := 0; i < 50; i++ {
		conn, err := http.Get("http://localhost" + addr + "/users")
		if err == nil {
			conn.Body.Close()
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	panic("server did not start")
}
