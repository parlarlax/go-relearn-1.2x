package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("=== 1. Match / MatchString ===")
	fmt.Println("match:", regexp.MustCompile(`^\d+$`).MatchString("12345"))
	fmt.Println("no match:", regexp.MustCompile(`^\d+$`).MatchString("12a45"))

	fmt.Println("\n=== 2. Compile / MustCompile ===")
	re := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
	fmt.Println("match email:", re.MatchString("user@example.com"))

	fmt.Println("\n=== 3. Find / FindString ===")
	re2 := regexp.MustCompile(`\d+`)
	fmt.Println("find:", re2.FindString("abc 123 def 456"))
	fmt.Println("find all:", re2.FindAllString("abc 123 def 456 ghi 789", -1))
	fmt.Println("find n=2:", re2.FindAllString("abc 123 def 456", 2))

	fmt.Println("\n=== 4. Find groups ===")
	re3 := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	match := re3.FindStringSubmatch("date: 2026-04-28")
	fmt.Println("full match:", match[0])
	fmt.Println("year:", match[1])
	fmt.Println("month:", match[2])
	fmt.Println("day:", match[3])

	fmt.Println("\n=== 5. Named groups ===")
	re4 := regexp.MustCompile(`(?P<first>\w+)\s+(?P<last>\w+)`)
	match2 := re4.FindStringSubmatch("John Doe")
	fmt.Println("named groups:", re4.SubexpNames())
	for i, name := range re4.SubexpNames() {
		if i > 0 && name != "" {
			fmt.Printf("  %s: %s\n", name, match2[i])
		}
	}

	fmt.Println("\n=== 6. Replace ===")
	re5 := regexp.MustCompile(`\bcat\b`)
	fmt.Println(re5.ReplaceAllString("the cat sat on the cat mat", "dog"))

	re6 := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
	result := re6.ReplaceAllString("email: user@example.com and admin@test.org", "[REDACTED]")
	fmt.Println("redact emails:", result)

	fmt.Println("\n=== 7. Replace with function ===")
	re7 := regexp.MustCompile(`\d+`)
	result2 := re7.ReplaceAllStringFunc("age: 25, score: 95", func(s string) string {
		return "***"
	})
	fmt.Println("mask numbers:", result2)

	fmt.Println("\n=== 8. Split ===")
	re8 := regexp.MustCompile(`\s*[,;]\s*`)
	parts := re8.Split("apple, banana;cherry ; date", -1)
	fmt.Println("split:", parts)

	fmt.Println("\n=== 9. String validation patterns ===")
	patterns := map[string]string{
		"email":    `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,
		"phone":    `^\d{3}-\d{3}-\d{4}$`,
		"hex":      `^#[0-9a-fA-F]{6}$`,
		"username": `^[a-zA-Z][a-zA-Z0-9_]{2,15}$`,
	}
	tests := map[string][]string{
		"email":    {"user@example.com", "bad email"},
		"phone":    {"123-456-7890", "12-34-56"},
		"hex":      {"#ff0000", "#gggggg"},
		"username": {"alice_123", "1badname"},
	}
	for name, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		for _, input := range tests[name] {
			fmt.Printf("  %s %q: %v\n", name, input, re.MatchString(input))
		}
	}
}
