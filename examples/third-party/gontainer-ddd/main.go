// gontainer + Domain-Driven Design example
//
// Structure:
//
//	domain/          ← entities + repository interfaces (pure Go, no DI dependency)
//	application/     ← use case services (depends on domain interfaces)
//	infrastructure/  ← concrete implementations: config, memory stores, middleware
//	interfaces/      ← HTTP handlers + server
//	main.go          ← gontainer wiring (equivalent to Spring @Configuration class)
//
// Java Spring Boot comparison:
//
//	SpringBootApplication.run()  → gontainer.Run()
//	@Configuration @Bean         → gontainer.NewFactory() in main.go
//	@Autowired constructor       → function parameters (auto-injected by type)
//	@Repository                  → infrastructure.MemoryUserStore (implements domain.UserStore)
//	@Service                     → application.UserService / OrderService
//	@RestController              → interfaces.UserHandler / OrderHandler
//	@Component Filter             → infrastructure.LoggingMiddleware / RecoveryMiddleware
//	@PreDestroy                   → return (service, cleanup) from factory
//	CommandLineRunner             → gontainer.NewEntrypoint()
//
// Note: gontainer matches services by exact type, not interface satisfaction.
// Factories return CONCRETE types (e.g. *infrastructure.MemoryUserStore).
// Application constructors still ACCEPT interfaces (e.g. domain.UserStore) —
// Go implicitly satisfies the interface, keeping the domain layer decoupled.
// This follows the Go idiom: "accept interfaces, return structs".
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

	"github.com/lax/go-relearn/examples/third-party/gontainer-ddd/application"
	"github.com/lax/go-relearn/examples/third-party/gontainer-ddd/infrastructure"
	"github.com/lax/go-relearn/examples/third-party/gontainer-ddd/interfaces"
)

func main() {
	fmt.Println("=== gontainer + Domain-Driven Design ===")
	fmt.Println()
	fmt.Println("Layers: domain → application → interfaces → infrastructure")
	fmt.Println("DI:     gontainer matches by exact type (concrete), not interface")
	fmt.Println()

	err := gontainer.Run(
		// ── Infrastructure Layer ──────────────────────────────
		// Spring Boot: @ConfigurationProperties
		gontainer.NewFactory(infrastructure.NewConfig),

		// Spring Boot: @Repository — return concrete type, gontainer matches by exact type
		gontainer.NewFactory(infrastructure.NewMemoryUserStore),
		gontainer.NewFactory(infrastructure.NewMemoryOrderStore),

		// Spring Boot: @Component implementing Filter interface
		// Registered as concrete types → gontainer.Multiple[HTTPMiddleware] collects them
		gontainer.NewFactory(infrastructure.NewLoggingMiddleware),
		gontainer.NewFactory(infrastructure.NewRecoveryMiddleware),

		// ── Application Layer ─────────────────────────────────
		// Spring Boot: @Service + @Autowired on constructor params
		// Note: NewUserService accepts domain.UserStore interface, but gontainer
		// injects the concrete *infrastructure.MemoryUserStore — Go satisfies the
		// interface implicitly, so the application layer stays decoupled from infrastructure.
		gontainer.NewFactory(func(users *infrastructure.MemoryUserStore) *application.UserService {
			return application.NewUserService(users)
		}),
		gontainer.NewFactory(func(
			orders *infrastructure.MemoryOrderStore,
			users *infrastructure.MemoryUserStore,
		) *application.OrderService {
			return application.NewOrderService(orders, users)
		}),

		// ── Interface Layer ───────────────────────────────────
		// Spring Boot: @RestController
		gontainer.NewFactory(interfaces.NewUserHandler),
		gontainer.NewFactory(interfaces.NewOrderHandler),

		// Spring Boot: @Bean Server + auto-collected Filters
		// Multiple[HTTPMiddleware] finds all factories returning types assignable to HTTPMiddleware
		gontainer.NewFactory(func(
			cfg *infrastructure.Config,
			mws gontainer.Multiple[infrastructure.HTTPMiddleware],
			uh *interfaces.UserHandler,
			oh *interfaces.OrderHandler,
		) (*interfaces.Server, func() error) {
			mux := http.NewServeMux()
			uh.RegisterRoutes(mux)
			oh.RegisterRoutes(mux)

			var handler http.Handler = mux
			for _, mw := range mws {
				handler = mw.Wrap(handler)
			}

			srv := interfaces.NewServer(cfg.Addr(), handler)
			return srv, func() error {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				return srv.Shutdown(ctx)
			}
		}),

		// ── Entrypoint ────────────────────────────────────────
		// Spring Boot: CommandLineRunner / ApplicationRunner
		gontainer.NewEntrypoint(func(srv *interfaces.Server, cfg *infrastructure.Config) {
			go srv.Start()
			waitForServer("http://localhost" + cfg.Addr())

			base := "http://localhost" + cfg.Addr()
			fmt.Println("--- HTTP Demo Requests ---")
			fmt.Println()

			fmt.Println("POST /users (create Alice)")
			post(base+"/users", `{"name":"Alice","email":"alice@test.com"}`)

			fmt.Println("POST /users (create Bob)")
			post(base+"/users", `{"name":"Bob","email":"bob@test.com"}`)

			fmt.Println("GET /users")
			get(base + "/users")

			fmt.Println("GET /users/1")
			get(base + "/users/1")

			fmt.Println("POST /orders (create order for Alice)")
			post(base+"/orders", `{"user_id":1,"item":"Go Book","qty":2}`)

			fmt.Println("POST /orders (create another order for Alice)")
			post(base+"/orders", `{"user_id":1,"item":"Coffee","qty":5}`)

			fmt.Println("GET /orders")
			get(base + "/orders")

			fmt.Println("GET /users/1/orders")
			get(base + "/users/1/orders")

			fmt.Println()
			fmt.Println("--- Cleanup (reverse order) ---")
		}),
	)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All services cleaned up. Done!")
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
		resp, err := http.Get(addr + "/users")
		if err == nil {
			resp.Body.Close()
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	panic("server did not start within 1 second")
}
