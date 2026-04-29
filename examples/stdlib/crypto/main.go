package main

import (
	"crypto/hkdf"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("=== 1. crypto/rand — secure random ===")
	bytes := make([]byte, 16)
	rand.Read(bytes)
	fmt.Println("random bytes:", hex.EncodeToString(bytes))

	n, _ := rand.Int(rand.Reader, big.NewInt(100))
	fmt.Println("random 0-99:", n)

	text := randomText(16)
	fmt.Println("random text:", text)

	fmt.Println("\n=== 2. sha256 ===")
	hash := sha256.Sum256([]byte("hello world"))
	fmt.Println("sha256:", hex.EncodeToString(hash[:]))

	h := sha256.New()
	h.Write([]byte("hello"))
	h.Write([]byte(" world"))
	fmt.Println("sha256 (streaming):", hex.EncodeToString(h.Sum(nil)))

	fmt.Println("\n=== 3. HKDF (Go 1.24+) ===")
	secret := []byte("my-secret-key")
	salt := make([]byte, 16)
	rand.Read(salt)
	info := "my-app-context"

	key1 := deriveKey(secret, salt, info, 32)
	key2 := deriveKey(secret, salt, info, 32)
	fmt.Println("derived key 1:", hex.EncodeToString(key1))
	fmt.Println("derived key 2:", hex.EncodeToString(key2))
	fmt.Println("keys match:", hex.EncodeToString(key1) == hex.EncodeToString(key2))
}

func randomText(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[n.Int64()]
	}
	return string(b)
}

func deriveKey(secret, salt []byte, info string, length int) []byte {
	key, err := hkdf.Key(sha256.New, secret, salt, info, length)
	if err != nil {
		panic(err)
	}
	return key
}
