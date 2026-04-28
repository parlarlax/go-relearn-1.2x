package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== strings ===")

	fmt.Println("\n--- Contains / HasPrefix / HasSuffix ---")
	s := "Hello, Go World"
	fmt.Println("Contains 'Go':", strings.Contains(s, "Go"))
	fmt.Println("HasPrefix 'Hello':", strings.HasPrefix(s, "Hello"))
	fmt.Println("HasSuffix 'World':", strings.HasSuffix(s, "World"))

	fmt.Println("\n--- Count / Index / LastIndex ---")
	fmt.Println("Count 'o':", strings.Count(s, "o"))
	fmt.Println("Index 'Go':", strings.Index(s, "Go"))
	fmt.Println("LastIndex 'o':", strings.LastIndex(s, "o"))

	fmt.Println("\n--- Replace / ReplaceAll ---")
	fmt.Println(strings.Replace(s, "o", "0", 1))
	fmt.Println(strings.ReplaceAll(s, "o", "0"))

	fmt.Println("\n--- Split / Join ---")
	parts := strings.Split("a,b,c,d", ",")
	fmt.Println("split:", parts)
	fmt.Println("join:", strings.Join(parts, "-"))

	fmt.Println("\n--- Trim / TrimSpace ---")
	fmt.Println(strings.TrimSpace("  hello  "))
	fmt.Println(strings.Trim("__hello__", "_"))

	fmt.Println("\n--- ToUpper / ToLower / Title ---")
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.ToLower("HELLO"))

	fmt.Println("\n--- Repeat ---")
	fmt.Println(strings.Repeat("ha", 3))

	fmt.Println("\n--- Builder (efficient concat) ---")
	var b strings.Builder
	for i := 0; i < 5; i++ {
		b.WriteString(fmt.Sprintf("item%d ", i))
	}
	fmt.Println("builder:", b.String())

	fmt.Println("\n--- Lines iterator (Go 1.24+) ---")
	text := "line1\nline2\nline3\n"
	for line := range strings.Lines(text) {
		fmt.Printf("  line: %q\n", line)
	}

	fmt.Println("\n--- SplitSeq iterator (Go 1.24+) ---")
	for part := range strings.SplitSeq("a-b-c-d", "-") {
		fmt.Printf("  part: %s\n", part)
	}

	fmt.Println("\n--- FieldsSeq iterator (Go 1.24+) ---")
	for field := range strings.FieldsSeq("  hello   world  go  ") {
		fmt.Printf("  field: %s\n", field)
	}

	fmt.Println("\n\n=== bytes ===")

	fmt.Println("\n--- Buffer (like strings.Builder for bytes) ---")
	var buf bytes.Buffer
	buf.WriteString("hello")
	buf.WriteByte(',')
	buf.WriteString(" world")
	fmt.Println("buffer:", buf.String())

	fmt.Println("\n--- bytes.Contains / Count / Index ---")
	data := []byte("Hello, Go World")
	fmt.Println("Contains 'Go':", bytes.Contains(data, []byte("Go")))
	fmt.Println("Count 'o':", bytes.Count(data, []byte("o")))

	fmt.Println("\n--- bytes.Split / Join ---")
	parts2 := bytes.Split([]byte("a,b,c"), []byte(","))
	fmt.Printf("split: %q\n", parts2)
	fmt.Printf("join: %s\n", bytes.Join(parts2, []byte("-")))

	fmt.Println("\n--- bytes.Trim ---")
	fmt.Printf("trim: %s\n", bytes.Trim([]byte("__hello__"), "_"))

	fmt.Println("\n--- Lines iterator (Go 1.24+) ---")
	for line := range bytes.Lines([]byte("foo\nbar\nbaz\n")) {
		fmt.Printf("  line: %q\n", line)
	}

	fmt.Println("\n--- SplitSeq iterator (Go 1.24+) ---")
	for part := range bytes.SplitSeq([]byte("x-y-z"), []byte("-")) {
		fmt.Printf("  part: %s\n", part)
	}
}
