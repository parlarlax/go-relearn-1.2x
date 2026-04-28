# Go Re-learning Lab (1.20 to 1.26) 🚀

Repository สำหรับทบทวนความรู้ภาษา Go และอัปเดตฟีเจอร์ใหม่ๆ ตั้งแต่เวอร์ชัน 1.20 เป็นต้นไป

## 🎯 Objectives

- เคาะสนิม Syntax พื้นฐาน และ Concurrency (Goroutines/Channels)
- ศึกษาฟีเจอร์ใหม่ที่สำคัญในแต่ละเวอร์ชัน (PGO, Iterators, Loop Semantics)
- สร้างโปรเจกต์ขนาดเล็กเพื่อทดลองใช้งานจริง

## 📅 Learning Roadmap

### Phase 1: The Basics (Refresher)

- [ ] Interface & Methods
- [ ] Concurrency Patterns (Select, Context, WaitGroup)
- [ ] Generics (revisit from 1.18+)

### Phase 2: What's New since 1.20?

- **Go 1.21:**
  - `min`, `max`, `clear` built-in functions
  - `log/slog` (Structured Logging - สำคัญมาก!)
  - `slices` และ `maps` package มาตรฐาน
- **Go 1.22:**
  - Loop semantics change (แก้ปัญหาตัวแปร loop ใน goroutine)
  - `net/http` routing ปรับปรุงใหม่ (Support method/wildcards)
- **Go 1.23:**
  - **Iterators** (Range-over-func - ฟีเจอร์ใหญ่!)
  - `unique` package
- **Go 1.24 - 1.26:**
  - Profile-Guided Optimization (PGO) improvements
  - การปรับปรุง Generic Alias และ Performance ต่างๆ

## 🧪 Experiments

(รายการโปรเจกต์ย่อยใน repo นี้)

- `/slog-demo`: ทดลองใช้ structured logging แทน log แบบเก่า
- `/http-mux`: เขียน API ด้วย router ใหม่ของ Go 1.22+
- `/iterators-lab`: ฝึกสร้าง custom iterator ด้วย Range-over-func

## 🛠 Prerequisites

- Go 1.26+ (ดาวน์โหลดได้ที่ [go.dev](https://go.dev/dl/))
- VS Code + Go Extension
