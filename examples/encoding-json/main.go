package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`                         // ไม่ serialize เลย
	Age       int       `json:"age,omitempty"`             // ข้ามถ้าเป็น zero value
	Bio       string    `json:"bio,omitzero"`              // ข้ามถ้าเป็น "" (Go 1.24+)
	CreatedAt time.Time `json:"created_at,omitzero"`       // ข้ามถ้าเป็น zero time (Go 1.24+)
	Role      string    `json:"role,omitempty"`
}

func main() {
	fmt.Println("=== 1. Marshal / Unmarshal ===")
	user := User{
		ID: 1, Name: "Alice", Email: "alice@test.com",
		Password: "secret123", Role: "admin",
	}
	data, _ := json.Marshal(user)
	fmt.Println(string(data))

	fmt.Println("\n=== 2. Pretty print ===")
	pretty, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(pretty))

	fmt.Println("\n=== 3. Unmarshal ===")
	input := `{"id":2,"name":"Bob","email":"bob@test.com","role":"user"}`
	var parsed User
	json.Unmarshal([]byte(input), &parsed)
	fmt.Printf("%+v\n", parsed)

	fmt.Println("\n=== 4. omitempty vs omitzero ===")
	empty := User{ID: 3, Name: "Charlie"}
	data, _ = json.Marshal(empty)
	fmt.Println("omitempty + omitzero:", string(data))

	fmt.Println("\n=== 5. json:\"-\" hides field ===")
	fmt.Println("Password is never in JSON output above")

	fmt.Println("\n=== 6. Dynamic JSON with map ===")
	var m map[string]any
	json.Unmarshal([]byte(`{"name":"test","count":42,"active":true}`), &m)
	fmt.Printf("map: %+v (type of count: %T)\n", m["count"], m["count"])

	fmt.Println("\n=== 7. Encoder / Decoder (streaming) ===")
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(map[string]string{"streaming": "works"})

	fmt.Println("\n=== 8. RawMessage (defer parsing) ===")
	raw := `{"meta":"data","nested":{"a":1}}`
	var obj struct {
		Meta    string          `json:"meta"`
		Nested  json.RawMessage `json:"nested"`
	}
	json.Unmarshal([]byte(raw), &obj)
	fmt.Printf("meta=%s, nested (raw)=%s\n", obj.Meta, string(obj.Nested))

	fmt.Println("\n=== 9. Custom MarshalJSON ===")
	event := Event{Title: "Go Meetup", Date: time.Date(2026, 6, 15, 18, 0, 0, 0, time.UTC)}
	data, _ = json.Marshal(event)
	fmt.Println(string(data))

	fmt.Println("\n=== 10. JSON validation with DisallowUnknownFields ===")
	dec := json.NewDecoder(nil)
	_ = dec
	fmt.Println("(use dec.DisallowUnknownFields() to reject unknown fields)")
}

type Event struct {
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

func (e Event) MarshalJSON() ([]byte, error) {
	type Alias Event
	return json.Marshal(&struct {
		Alias
		Date string `json:"date"`
	}{
		Alias: Alias(e),
		Date:  e.Date.Format("2006-01-02"),
	})
}
