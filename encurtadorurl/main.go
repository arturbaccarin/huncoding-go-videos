// https://youtu.be/VoykQd8Q5iU
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"sync"
)

var (
	urlStore   = make(map[string]string)
	secretKey  = []byte("aaaaaaaabbbbbbbbccccccccdddddddd")
	letterRune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	mu         sync.Mutex
)

func shortenURL(w http.ResponseWriter, r *http.Request) {
	originalURL := r.URL.Query().Get("url")
	if originalURL == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	encryptedURL := encryptURL(originalURL)
	shortID := generateShortID()

	mu.Lock()
	urlStore[shortID] = encryptedURL
	mu.Unlock()

	shortURL := fmt.Sprintf("http://localhost:8080/%s", shortID)
	fmt.Fprintf(w, "Short URL: %s\n", shortURL)
}

func encryptURL(originalURL string) string {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	plainText := []byte(originalURL)
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	iv := cipherText[:aes.BlockSize]

	if _, err := rand.Read(iv); err != nil {
		log.Fatal(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return hex.EncodeToString(cipherText)
}

func generateShortID() string {
	b := make([]rune, 6)

	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterRune))))
		if err != nil {
			log.Fatal(err)
		}

		b[i] = letterRune[num.Int64()]
	}

	return string(b)
}

func main() {
	fmt.Println(generateShortID())
}
