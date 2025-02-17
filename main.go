package main

import (
	"context"
	"log/slog"
	"net/http"
	"time"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
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



// API request function with OpenTelemetry tracing
func makeAPIRequest(ctx context.Context, url string) error {
	tracer := otel.Tracer("project-vending-module")

	ctx, span := tracer.Start(ctx, "API Request")
	defer span.End()

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)

	if err != nil {
		span.RecordError(err)
		fmt.Println("API request failed:", err)
		return err
	}
	defer resp.Body.Close()

	span.SetAttributes("http.status_code", resp.StatusCode)
	fmt.Println("API request successful:", resp.StatusCode)

	return nil
}


func main() {
	ctx := context.Background()
	_ = makeAPIRequest(ctx, "https://example.com/api")
}
