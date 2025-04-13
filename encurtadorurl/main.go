// https://youtu.be/VoykQd8Q5iU
// commit of the day
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
	"strings"
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

	if !(strings.HasPrefix(originalURL, "http://") || strings.HasPrefix(originalURL, "https://")) {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
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

func decryptURL(encryptedURL string) string {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	cipherText, err := hex.DecodeString(encryptedURL)
	if err != nil {
		log.Fatal(err)
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortID := r.URL.Path[1:]

	mu.Lock()
	encryptedURL, ok := urlStore[shortID]
	mu.Unlock()

	if !ok {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	decryptedURL := decryptURL(encryptedURL)
	http.Redirect(w, r, decryptedURL, http.StatusSeeOther)
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
	http.HandleFunc("/shorten", shortenURL)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
