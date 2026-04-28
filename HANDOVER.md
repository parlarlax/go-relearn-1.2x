# Project Handoff / Progress Tracker

## Current Status

- **Last Active:** Phase 1 + Phase 2 สร้างตัวอย่างครบแล้ว
- **Current Branch:** `main`
- **Go Version:** 1.26.1

## What's Working

- [x] Setup `go.mod` และโครงสร้างโฟลเดอร์ (Idiomatic Go style)
- [x] Phase 1.1: Interface & Methods (6 ตัวอย่าง)
- [x] Phase 1.2: Concurrency Patterns (6 ตัวอย่าง)
- [x] Phase 1.3: Generics (4 ตัวอย่าง)
- [x] Phase 2.1: slog-demo (3 ตัวอย่าง)
- [x] Phase 2.2: http-mux (Go 1.22+ new router)
- [x] Phase 2.3: iterators-lab (basic + tree iterator)
- [x] `go vet ./...` ผ่านหมด

## Next Steps

1. ลองรันแต่ละ experiment ด้วย `go run ./experiments/...`
2. ศึกษา `unique` package (Go 1.23)
3. ทดลอง PGO (Profile-Guided Optimization) ในเครื่อง local
4. เพิ่ม unit tests ให้แต่ละ package ใน `basics/`
5. ลองใช้ `slices` และ `maps` standard package (Go 1.21+)

## Quick Tips

- ใช้ `go vet ./...` ตรวจสอบทุกครั้งก่อน commit
- ดูตัวอย่าง Iterator ใน `experiments/iterators-lab/basic/`
- http-mux ต้องรันแล้วใช้ curl ทดสอบ: `curl localhost:8080/users/42`

## Blockers

- (ยังไม่มี)
