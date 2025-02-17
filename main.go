package main

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

// Example API request function with structured logging
func makeAPIRequest(ctx context.Context, url string) error {
	logger := slog.Default()

	logger.Info("Starting API request", "url", url)

	// Simulate an API call
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)

	if err != nil {
		logger.Error("API request failed", "url", url, "error", err)
		return err
	}
	defer resp.Body.Close()

	logger.Info("API request completed", "url", url, "status", resp.StatusCode)

	return nil
}

func main() {
	ctx := context.Background()
	_ = makeAPIRequest(ctx, "https://example.com/api")
}
