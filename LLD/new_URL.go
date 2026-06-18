package main

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type URLShortener struct {
	longToShort map[string]string
	shortToLong map[string]string
	baseURL     string
}

func NewURLShortener() *URLShortener {

	rand.Seed(time.Now().UnixNano())

	return &URLShortener{
		longToShort: make(map[string]string),
		shortToLong: make(map[string]string),
		baseURL:     "http://tinyurl.com/",
	}
}

func (u *URLShortener) Shorten(longURL string) string {

	// If already shortened, return existing value
	if shortCode, exists := u.longToShort[longURL]; exists {
		return u.baseURL + shortCode
	}

	shortCode := u.generateCode(6)

	// Collision handling
	for {
		if _, exists := u.shortToLong[shortCode]; !exists {
			break
		}
		shortCode = u.generateCode(6)
	}

	u.longToShort[longURL] = shortCode
	u.shortToLong[shortCode] = longURL

	return u.baseURL + shortCode
}

func (u *URLShortener) Expand(shortURL string) string {

	shortCode := shortURL[len(u.baseURL):]

	if longURL, exists := u.shortToLong[shortCode]; exists {
		return longURL
	}

	return ""
}

func (u *URLShortener) generateCode(length int) string {

	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

func main() {

	shortener := NewURLShortener()

	longURL := "https://leetcode.com/problems/lru-cache"

	shortURL := shortener.Shorten(longURL)

	fmt.Println("Short URL:", shortURL)

	originalURL := shortener.Expand(shortURL)

	fmt.Println("Original URL:", originalURL)
}