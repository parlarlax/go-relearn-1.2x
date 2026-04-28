# 🏁 Project Handoff / Progress Tracker

ไฟล์นี้ใช้บันทึกสถานะล่าสุดว่าทำถึงไหนแล้ว และต้อง "ไปต่อ" ที่จุดไหน เพื่อให้การกลับมาเขียนต่อทำได้ลื่นไหลที่สุด

## 📍 Current Status (สถานะปัจจุบัน)

- **Last Active:** [ใส่สถานะล่าสุด เช่น กำลังทดลองเรื่อง Iterators ใน Go 1.23]
- **Current Branch:** `main` (หรือชื่อฟีเจอร์ที่ทำค้างไว้)
- **Latest Milestone:** ติดตั้ง Go 1.26 และเซ็ตอัพโปรเจกต์พื้นฐานสำเร็จ

## 🏗 What's Working (สิ่งที่ทำเสร็จแล้ว)

- [ ] Setup `go.mod` และโครงสร้างโฟลเดอร์
- [ ] ทดสอบ `log/slog` เบื้องต้น
- [ ] ลองใช้ `net/http` router ตัวใหม่ (v1.22+)

## 🚧 Work in Progress (สิ่งที่ค้างอยู่)

- [ ] กำลังแก้บัคในส่วนของ `custom iterator` ในโฟลเดอร์ `/experiments/iter`
- [ ] ยังทำ Unit Test ของส่วน `Service` ไม่เสร็จ

## 📝 Next Steps (สิ่งที่ต้องทำต่อ)

1. ไปที่ไฟล์ `internal/handler/user_test.go` แล้วเขียน Test Case เพิ่ม
2. ศึกษาเรื่อง **PGO (Profile-Guided Optimization)** ว่าจะลองรันในเครื่อง local อย่างไร
3. อัปเดต `README.md` ในส่วนของ Roadmap เมื่อทำหัวข้อนั้นๆ จบ

## 💡 Quick Tips & Reminders

- อย่าลืมใช้คำสั่ง `go mod tidy` ทุกครั้งหลังเพิ่ม package ใหม่
- ดูตัวอย่างการเขียน Iterator ได้ใน [Go 1.23 Release Notes](https://go.dev)
- ติดปัญหาตรงไหนให้โน้ตไว้ในส่วน **Blockers** ด้านล่าง

## 🛑 Blockers (ปัญหาที่พบ)

- (ตัวอย่าง) ยังไม่ค่อยเข้าใจเรื่อง `unique` package ว่าใช้จริงตอนไหนดี -> ต้องหาบทความอ่านเพิ่ม
