package main

import (
	"context"
	"log/slog"
	"os"
	"testing"
)

// TestLoggerHandler only logs when a test fails
type TestLoggerHandler struct {
	slog.Handler
	t *testing.T
}

func (h *TestLoggerHandler) Handle(ctx context.Context, r slog.Record) error {
	// Only log if the test has failed
	if h.t.Failed() {
		return h.Handler.Handle(ctx, r)
	}
	return nil
}

// SetTestLogger sets the global slog logger to a test-aware handler
func SetTestLogger(t *testing.T) {
	baseHandler := slog.NewTextHandler(os.Stdout, nil) // You can switch to JSON if needed
	testHandler := &TestLoggerHandler{Handler: baseHandler, t: t}
	slog.SetDefault(slog.New(testHandler))
}
