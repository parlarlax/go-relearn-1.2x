package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	_ "modernc.org/sqlite"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		slog.Error("open db", "error", err)
		showPatterns()
		return
	}
	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	realDemo(db)
}

// Java comparison:
//   JdbcTemplate → sql.DB (raw SQL)
//   @Transactional → db.Begin() / tx.Commit() / tx.Rollback()
//   PreparedStatement → db.Prepare() / stmt.QueryRow()
//   ResultSet → rows.Next() / rows.Scan()
//   Spring Data JPA → no direct equivalent; consider GORM or sqlc

func realDemo(db *sql.DB) {
	fmt.Println("=== 1. Create table ===")
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id   INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE
		)
	`)
	if err != nil {
		slog.Error("create table", "error", err)
		return
	}
	fmt.Println("table created")

	fmt.Println("\n=== 2. Insert (Exec) ===")
	result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Alice", "alice@test.com")
	if err != nil {
		slog.Error("insert", "error", err)
		return
	}
	id, _ := result.LastInsertId()
	fmt.Println("inserted id:", id)

	db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Bob", "bob@test.com")
	db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Charlie", "charlie@test.com")

	fmt.Println("\n=== 3. Query single row ===")
	var u User
	err = db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", 1).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		fmt.Println("query row:", err)
	} else {
		data, _ := json.Marshal(u)
		fmt.Println("found:", string(data))
	}

	fmt.Println("\n=== 4. Query multiple rows ===")
	rows, err := db.Query("SELECT id, name, email FROM users ORDER BY id")
	if err != nil {
		slog.Error("query", "error", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name, &user.Email)
		fmt.Printf("  [%d] %s (%s)\n", user.ID, user.Name, user.Email)
	}

	fmt.Println("\n=== 5. Update ===")
	res, _ := db.Exec("UPDATE users SET name = ? WHERE id = ?", "Alice V2", 1)
	affected, _ := res.RowsAffected()
	fmt.Println("updated rows:", affected)

	fmt.Println("\n=== 6. Delete ===")
	res, _ = db.Exec("DELETE FROM users WHERE id = ?", 3)
	affected, _ = res.RowsAffected()
	fmt.Println("deleted rows:", affected)

	fmt.Println("\n=== 7. Transaction ===")
	tx, err := db.Begin()
	if err != nil {
		slog.Error("begin tx", "error", err)
		return
	}
	tx.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Dave", "dave@test.com")
	tx.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Eve", "eve@test.com")
	tx.Commit()
	fmt.Println("transaction committed")

	fmt.Println("\n=== 8. Prepared statement ===")
	stmt, err := db.Prepare("SELECT name, email FROM users WHERE id = ?")
	if err != nil {
		slog.Error("prepare", "error", err)
		return
	}
	defer stmt.Close()
	var name, email string
	stmt.QueryRow(1).Scan(&name, &email)
	fmt.Printf("prepared: %s (%s)\n", name, email)

	fmt.Println("\n=== 9. Context pattern ===")
	fmt.Println("  db.QueryContext(ctx, \"SELECT ...\", args...)  — timeout/cancel")
	fmt.Println("  db.ExecContext(ctx, \"INSERT ...\", args...)    — timeout/cancel")
}

func showPatterns() {
	fmt.Println()
	fmt.Println("=== database/sql patterns (no DB connection) ===")

	fmt.Println("\n--- Open + Connection Pool ---")
	fmt.Println("  db, err := sql.Open(\"postgres\", connStr)")
	fmt.Println("  db.SetMaxOpenConns(10)")
	fmt.Println("  db.SetMaxIdleConns(5)")
	fmt.Println("  db.SetConnMaxLifetime(30 * time.Minute)")
	fmt.Println("  defer db.Close()")

	fmt.Println("\n--- Insert (Exec) ---")
	fmt.Println("  result, err := db.Exec(\"INSERT INTO users (name, email) VALUES (?, ?)\", name, email)")
	fmt.Println("  id, _ := result.LastInsertId()")

	fmt.Println("\n--- Query single row ---")
	fmt.Println("  var u User")
	fmt.Println("  err := db.QueryRow(\"SELECT id, name FROM users WHERE id = ?\", 1).Scan(&u.ID, &u.Name)")

	fmt.Println("\n--- Query multiple rows ---")
	fmt.Println("  rows, err := db.Query(\"SELECT id, name FROM users ORDER BY id\")")
	fmt.Println("  defer rows.Close()")
	fmt.Println("  for rows.Next() { rows.Scan(&id, &name) }")

	fmt.Println("\n--- Transaction ---")
	fmt.Println("  tx, _ := db.Begin()")
	fmt.Println("  tx.Exec(\"INSERT ...\", ...)")
	fmt.Println("  tx.Commit()  // or tx.Rollback()")

	fmt.Println("\n--- Prepared statement ---")
	fmt.Println("  stmt, _ := db.Prepare(\"SELECT name FROM users WHERE id = ?\")")
	fmt.Println("  defer stmt.Close()")
	fmt.Println("  stmt.QueryRow(1).Scan(&name)")

	fmt.Println("\n--- With Context (timeout/cancel) ---")
	fmt.Println("  ctx, cancel := context.WithTimeout(ctx, 5*time.Second)")
	fmt.Println("  db.QueryContext(ctx, \"SELECT ...\")")
	fmt.Println("  db.ExecContext(ctx, \"INSERT ...\")")

	fmt.Println("\n--- Supported drivers ---")
	fmt.Println("  github.com/lib/pq           — PostgreSQL")
	fmt.Println("  go-sql-driver/mysql          — MySQL")
	fmt.Println("  modernc.org/sqlite           — SQLite3 (pure-Go, no CGO)")
	fmt.Println("  github.com/jackc/pgx        — PostgreSQL (modern)")
}
