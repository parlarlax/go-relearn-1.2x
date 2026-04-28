# Go Re-learning Lab (1.20 to 1.26) 🚀

Repository สำหรับทบทวนความรู้ภาษา Go และอัปเดตฟีเจอร์ใหม่ๆ ตั้งแต่เวอร์ชัน 1.20 เป็นต้นไป

## 🎯 Objectives

- เคาะสนิม Syntax พื้นฐาน และ Concurrency (Goroutines/Channels)
- ศึกษาฟีเจอร์ใหม่ที่สำคัญในแต่ละเวอร์ชัน (PGO, Iterators, Loop Semantics)
- สร้างโปรเจกต์ขนาดเล็กเพื่อทดลองใช้งานจริง

## 📁 Project Structure

```
cmd/basics/              — Entry point สำหรับรันตัวอย่าง Phase 1
basics/
  interfaces/            — Phase 1.1: Interface & Methods
  concurrency/           — Phase 1.2: Concurrency Patterns
  generics/              — Phase 1.3: Generics
experiments/
  slog-demo/             — Phase 2.1: log/slog structured logging
    text-handler/        — TextHandler example
    json-handler/        — JSONHandler + slog.Group example
    custom-handler/      — Custom handler wrapper example
  http-mux/              — Phase 2.2: Go 1.22+ new net/http router
  iterators-lab/         — Phase 2.3: Go 1.23 range-over-func iterators
    basic/               — Backward, Filter, Enumerate iterators
    tree/                — BinaryTree iterator with iter.Seq
```

## 📅 Learning Roadmap

### Phase 1: The Basics (Refresher)

- [x] Interface & Methods (`basics/interfaces/`)
- [x] Concurrency Patterns — Select, Context, WaitGroup (`basics/concurrency/`)
- [x] Generics (`basics/generics/`)

### Phase 2: What's New since 1.20?

- **Go 1.21:**
  - `min`, `max`, `clear` built-in functions
  - `log/slog` (Structured Logging)
  - `slices` และ `maps` package มาตรฐาน
- **Go 1.22:**
  - Loop semantics change
  - `net/http` routing ปรับปรุงใหม่ (`experiments/http-mux/`)
- **Go 1.23:**
  - **Iterators** (Range-over-func) (`experiments/iterators-lab/`)
  - `unique` package
- **Go 1.24 - 1.26:**
  - Profile-Guided Optimization (PGO) improvements
  - Generic Alias และ Performance ต่างๆ

## 🏃 How to Run

```bash
# Phase 1 — ผ่าน cmd entry point
go run ./cmd/basics interfaces

# Phase 2 — รันแต่ละ experiment ตรงๆ
go run ./experiments/slog-demo/text-handler
go run ./experiments/slog-demo/json-handler
go run ./experiments/slog-demo/custom-handler
go run ./experiments/http-mux
go run ./experiments/iterators-lab/basic
go run ./experiments/iterators-lab/tree
```

## 🛠 Prerequisites

- Go 1.26+ (ดาวน์โหลดได้ที่ [go.dev](https://go.dev/dl/))
- VS Code + Go Extension
