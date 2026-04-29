package main

import (
	"errors"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/NVIDIA/gontainer/v2"
)

// ============================================================
// Domain types — Spring Boot equivalents in comments
// ============================================================

// @Component Config — similar to @ConfigurationProperties
type Config struct {
	AppName string
	Port    int
	Debug   bool
}

// @Component Logger — similar to injecting a SLF4J Logger
type Logger struct {
	handler slog.Handler
}

func (l *Logger) Info(msg string, args ...any) {
	if !l.handler.Enabled(nil, slog.LevelInfo) {
		return
	}
	l.handler.Handle(nil, slog.Record{
		Message: fmt.Sprintf(msg, args...),
		Level:   slog.LevelInfo,
		Time:    time.Now(),
	})
}

// @Repository Database — similar to @Bean DataSource
type Database struct {
	dsn string
}

func (d *Database) Query(table string) string {
	return fmt.Sprintf("SELECT * FROM %s (connected: %s)", table, d.dsn)
}

func (d *Database) Close() error {
	fmt.Printf("  [Database] closing connection: %s\n", d.dsn)
	return nil
}

// @Service UserRepository — similar to Spring Data Repository
type UserRepository struct {
	db *Database
}

func (r *UserRepository) FindAll() string {
	return r.db.Query("users")
}

// @Service EmailService — similar to @Service with @Value
type EmailService struct {
	smtp string
	log  *Logger
}

func (s *EmailService) Send(to, body string) {
	s.log.Info("sending email to %s: %s", to, body)
}

// @Middleware interface — similar to HandlerInterceptor
type Middleware interface {
	Name() string
}

// @Component AuthMiddleware — implements Middleware
type AuthMiddleware struct{}

func (m *AuthMiddleware) Name() string { return "Auth" }

// @Component LogMiddleware — implements Middleware
type LogMiddleware struct{}

func (m *LogMiddleware) Name() string { return "Log" }

// @Component MetricsService — optional dependency, may not be registered
type MetricsService struct {
	endpoint string
}

// @Service Router — depends on multiple Middleware + optional MetricsService
type Router struct {
	middlewares []Middleware
	metrics     *MetricsService
	log         *Logger
}

func (r *Router) Handle(path string) {
	mwNames := make([]string, len(r.middlewares))
	for i, m := range r.middlewares {
		mwNames[i] = m.Name()
	}
	metricsInfo := "disabled"
	if r.metrics != nil {
		metricsInfo = "enabled (" + r.metrics.endpoint + ")"
	}
	r.log.Info("route %q | middleware: [%s] | metrics: %s", path, strings.Join(mwNames, ", "), metricsInfo)
}

// Transaction — transient (new instance each call), similar to prototype scope
type Transaction struct {
	ID int
}

// ============================================================
// Factory functions — @Bean methods equivalent
// ============================================================

func newConfig() *Config {
	return &Config{AppName: "go-relearn-demo", Port: 8080, Debug: true}
}

func newLogger(cfg *Config) *Logger {
	var handler slog.Handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	return &Logger{handler: handler}
}

func newDatabase(cfg *Config) (*Database, func() error) {
	db := &Database{dsn: fmt.Sprintf("sqlite://:%d/%s", cfg.Port, cfg.AppName)}
	return db, db.Close
}

func newUserRepository(db *Database) *UserRepository {
	return &UserRepository{db: db}
}

func newEmailService(cfg *Config, log *Logger) *EmailService {
	return &EmailService{smtp: fmt.Sprintf("smtp://%s", cfg.AppName), log: log}
}

func newAuthMiddleware() *AuthMiddleware { return &AuthMiddleware{} }
func newLogMiddleware() *LogMiddleware   { return &LogMiddleware{} }

func newMetricsService() *MetricsService {
	return &MetricsService{endpoint: "http://localhost:9090/metrics"}
}

func newRouter(
	middlewares gontainer.Multiple[Middleware],
	metrics gontainer.Optional[*MetricsService],
	log *Logger,
) *Router {
	r := &Router{log: log, metrics: metrics.Get()}
	for _, mw := range middlewares {
		r.middlewares = append(r.middlewares, mw)
	}
	return r
}

func newTransactionFactory() func() *Transaction {
	return func() *Transaction {
		return &Transaction{ID: rand.Int()}
	}
}

// ============================================================
// Annotations — metadata without starting container
// ============================================================

type serviceInfo struct {
	Name string
	Desc string
}

// ============================================================
// main — gontainer.Run is similar to SpringApplication.run()
// ============================================================

func main() {
	fmt.Println("=== NVIDIA gontainer — Dependency Injection for Go ===")
	fmt.Println()
	fmt.Println("Java Spring Boot equivalent:")
	fmt.Println("  @Component / @Service  → gontainer.NewFactory()")
	fmt.Println("  @Bean                   → gontainer.NewFactory()")
	fmt.Println("  @Autowired              → function parameters (auto-injected)")
	fmt.Println("  @Qualifier              → named types (type UsersDB *sql.DB)")
	fmt.Println("  @Optional               → gontainer.Optional[T]")
	fmt.Println("  List<Middleware>        → gontainer.Multiple[T]")
	fmt.Println("  @Scope('prototype')     → factory returning func() *T")
	fmt.Println("  @PreDestroy             → return (service, cleanup) from factory")
	fmt.Println("  DisposableBean.destroy() → cleanup func() error")
	fmt.Println("  ApplicationContext      → *gontainer.Resolver")
	fmt.Println("  SpringApplication.run() → gontainer.Run()")
	fmt.Println()

	fmt.Println("--- 1. Basic: Factory + Auto-injection ---")
	runBasicDemo()

	fmt.Println()
	fmt.Println("--- 2. Cleanup: Lifecycle management (reverse order) ---")
	runCleanupDemo()

	fmt.Println()
	fmt.Println("--- 3. Optional + Multiple dependencies ---")
	runOptionalMultipleDemo()

	fmt.Println()
	fmt.Println("--- 4. Transient services (prototype scope) ---")
	runTransientDemo()

	fmt.Println()
	fmt.Println("--- 5. Annotations (pre-run metadata) ---")
	runAnnotationsDemo()

	fmt.Println()
	fmt.Println("--- 6. Error handling ---")
	runErrorDemo()
}

func runBasicDemo() {
	err := gontainer.Run(
		gontainer.NewFactory(newConfig),
		gontainer.NewFactory(newLogger),
		gontainer.NewFactory(newDatabase),
		gontainer.NewFactory(newUserRepository),
		gontainer.NewFactory(newEmailService),
		gontainer.NewEntrypoint(func(repo *UserRepository, email *EmailService, cfg *Config) {
			fmt.Println("  config:", cfg.AppName, "port:", cfg.Port)
			fmt.Println("  users:", repo.FindAll())
			email.Send("alice@test.com", "Welcome!")
		}),
	)
	if err != nil {
		fmt.Println("  error:", err)
	}
}

func runCleanupDemo() {
	err := gontainer.Run(
		gontainer.NewFactory(newConfig),
		gontainer.NewFactory(newLogger),
		gontainer.NewFactory(newDatabase),
		gontainer.NewEntrypoint(func(db *Database) {
			fmt.Println("  using db:", db.Query("orders"))
		}),
	)
	if err != nil {
		fmt.Println("  error:", err)
	}
	fmt.Println("  (Database cleanup ran automatically in reverse order)")
}

func runOptionalMultipleDemo() {
	err := gontainer.Run(
		gontainer.NewFactory(newConfig),
		gontainer.NewFactory(newLogger),
		gontainer.NewFactory(newAuthMiddleware),
		gontainer.NewFactory(newLogMiddleware),
		gontainer.NewFactory(newMetricsService),
		gontainer.NewFactory(newRouter),
		gontainer.NewEntrypoint(func(router *Router) {
			router.Handle("/api/users")
			router.Handle("/api/orders")
		}),
	)
	if err != nil {
		fmt.Println("  error:", err)
	}
}

func runTransientDemo() {
	err := gontainer.Run(
		gontainer.NewFactory(newTransactionFactory),
		gontainer.NewEntrypoint(func(newTx func() *Transaction) {
			tx1 := newTx()
			tx2 := newTx()
			tx3 := newTx()
			fmt.Printf("  tx1.ID=%d, tx2.ID=%d, tx3.ID=%d (all different)\n", tx1.ID, tx2.ID, tx3.ID)
		}),
	)
	if err != nil {
		fmt.Println("  error:", err)
	}
}

func runAnnotationsDemo() {
	cfgFactory := gontainer.NewFactory(
		newConfig,
		gontainer.WithAnnotation(serviceInfo{Name: "config", Desc: "Application configuration"}),
	)
	dbFactory := gontainer.NewFactory(
		newDatabase,
		gontainer.WithAnnotation(serviceInfo{Name: "database", Desc: "SQLite database connection"}),
	)

	fmt.Println("  Registered services (inspected without starting container):")
	for _, f := range []*gontainer.Factory{cfgFactory, dbFactory} {
		for _, a := range f.Annotations() {
			if info, ok := a.(serviceInfo); ok {
				fmt.Printf("    - %-12s %s\n", info.Name+":", info.Desc)
			}
		}
	}
}

func runErrorDemo() {
	fmt.Println("  Trying container with missing dependency:")
	err := gontainer.Run(
		gontainer.NewEntrypoint(func(db *Database) {
			fmt.Println("this should not run")
		}),
	)
	if err != nil {
		if errors.Is(err, gontainer.ErrDependencyNotResolved) {
			fmt.Println("  caught ErrDependencyNotResolved ✓")
		}
		fmt.Printf("  full error: %v\n", err)
	}

	fmt.Println()
	fmt.Println("  Trying circular dependency:")
	type A struct{ x int }
	type B struct{ x int }
	err = gontainer.Run(
		gontainer.NewFactory(func(b *B) *A { return &A{} }),
		gontainer.NewFactory(func(a *A) *B { return &B{} }),
		gontainer.NewEntrypoint(func(a *A) { fmt.Println("should not reach here") }),
	)
	if err != nil {
		if errors.Is(err, gontainer.ErrCircularDependency) {
			fmt.Println("  caught ErrCircularDependency ✓")
		}
		fmt.Printf("  full error: %v\n", err)
	}
}
