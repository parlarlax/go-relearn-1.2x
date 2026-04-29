package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("=== 1. ReadFile / WriteFile ===")
	data := []byte("Hello from Go!")
	os.WriteFile("/tmp/go-example.txt", data, 0644)
	content, _ := os.ReadFile("/tmp/go-example.txt")
	fmt.Println("read:", string(content))

	fmt.Println("\n=== 2. MkdirAll / RemoveAll ===")
	os.MkdirAll("/tmp/go-example/a/b/c", 0755)
	os.WriteFile("/tmp/go-example/a/b/c/file.txt", []byte("nested"), 0644)
	os.RemoveAll("/tmp/go-example")
	fmt.Println("created and removed nested dirs")

	fmt.Println("\n=== 3. Environment variables ===")
	os.Setenv("MY_APP_MODE", "debug")
	fmt.Println("MY_APP_MODE:", os.Getenv("MY_APP_MODE"))
	fmt.Println("HOME:", os.Getenv("HOME"))
	fmt.Println("MISSING:", os.Getenv("MISSING_XYZ"))
	for _, e := range os.Environ() {
		if len(e) > 10 {
			continue
		}
		fmt.Println("  env:", e)
	}

	fmt.Println("\n=== 4. Args ===")
	fmt.Println("program:", os.Args[0])
	fmt.Println("args:", os.Args[1:])

	fmt.Println("\n=== 5. Stdin / Stdout / Stderr ===")
	fmt.Fprintf(os.Stdout, "to stdout\n")
	fmt.Fprintf(os.Stderr, "to stderr\n")

	fmt.Println("\n=== 6. Getwd / Chdir ===")
	wd, _ := os.Getwd()
	fmt.Println("cwd:", wd)

	fmt.Println("\n=== 7. Stat / IsDir ===")
	info, err := os.Stat("main.go")
	if err != nil {
		fmt.Println("stat error:", err)
	} else {
		fmt.Printf("name=%s size=%d dir=%v mode=%s\n", info.Name(), info.Size(), info.IsDir(), info.Mode())
	}

	fmt.Println("\n=== 8. ReadDir ===")
	entries, _ := os.ReadDir(".")
	for _, e := range entries {
		suffix := ""
		if e.IsDir() {
			suffix = "/"
		}
		fmt.Printf("  %s%s\n", e.Name(), suffix)
	}

	fmt.Println("\n=== 9. TempDir / TempFile ===")
	tmpDir, _ := os.MkdirTemp("", "go-example-*")
	fmt.Println("temp dir:", tmpDir)
	tmpFile, _ := os.CreateTemp("", "go-example-*.txt")
	fmt.Println("temp file:", tmpFile.Name())
	tmpFile.Close()
	os.RemoveAll(tmpDir)
	os.Remove(tmpFile.Name())

	fmt.Println("\n=== 10. os.Root sandbox (Go 1.24+) ===")
	root, err := os.OpenRoot(".")
	if err != nil {
		fmt.Println("open root error:", err)
		return
	}
	f, err := root.Open("main.go")
	if err != nil {
		fmt.Println("open safe:", err)
	} else {
		fmt.Println("os.Root: opened main.go safely")
		f.Close()
	}
	_, err = root.Open("../../../etc/passwd")
	if err != nil {
		fmt.Println("os.Root: blocked escape:", err)
	}

	fmt.Println("\n=== 11. Signal handling ===")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("signal handler registered (Ctrl+C to test)")
}
