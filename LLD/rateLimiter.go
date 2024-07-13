package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// RateLimiter struct to hold the rate limiting parameters and state
type RateLimiter struct {
	mu            sync.Mutex
	requests      map[int64]int
	limit         int
	windowSize    int64
	cleanupTicker *time.Ticker
}

// NewRateLimiter creates a new RateLimiter
func NewRateLimiter(limit int, windowSize int64) *RateLimiter {
	rl := &RateLimiter{
		requests:      make(map[int64]int),
		limit:         limit,
		windowSize:    windowSize,
		cleanupTicker: time.NewTicker(time.Duration(windowSize) * time.Second),
	}

	go rl.cleanupExpiredEntries()

	return rl
}

// Allow checks if a request is allowed
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now().Unix()

	// Clean up old entries
	for timestamp := range rl.requests {
		if now-timestamp >= rl.windowSize {
			delete(rl.requests, timestamp)
		}
	}

	// Sum the requests in the current window
	requestCount := 0
	for timestamp, count := range rl.requests {
		if now-timestamp < rl.windowSize {
			requestCount += count
		}
	}

	if requestCount >= rl.limit {
		return false
	}

	// Record the new request
	rl.requests[now]++
	return true
}

// cleanupExpiredEntries periodically removes expired entries from the requests map
func (rl *RateLimiter) cleanupExpiredEntries() {
	for range rl.cleanupTicker.C {
		rl.mu.Lock()
		now := time.Now().Unix()
		for timestamp := range rl.requests {
			if now-timestamp >= rl.windowSize {
				delete(rl.requests, timestamp)
			}
		}
		rl.mu.Unlock()
	}
}

// RateLimiterMiddleware applies rate limiting to HTTP requests
func RateLimiterMiddleware(next http.Handler, rateLimiter *RateLimiter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rateLimiter.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// HelloHandler is a simple HTTP handler for demonstration
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func main() {
	// Create a new rate limiter allowing 5 requests per 10 seconds
	rateLimiter := NewRateLimiter(5, 10)

	// Create a new HTTP server
	mux := http.NewServeMux()
	mux.Handle("/", RateLimiterMiddleware(http.HandlerFunc(HelloHandler), rateLimiter))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server is running on port 8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
