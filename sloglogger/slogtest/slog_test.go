package main

import (
	"log/slog"
	"testing"
)

func TestInitial(t *testing.T) {
	t.Run("TestExampleFail", func(t *testing.T) {
		SetTestLogger(t) // ✅ Each subtest gets its own logger
		TestExampleFail(t)
	})

	t.Run("TestExamplePass", func(t *testing.T) {
		SetTestLogger(t) // ✅ Each subtest gets its own logger
		TestExamplePass(t)
	})
}

// TestExample Example test using slog without explicit injection
func TestExampleFail(t *testing.T) {
	slog.Info("This will only be logged if the test fails")
	t.Fail() // Simulating a failure
	slog.Info("This log will now appear since the test has failed")
}

func TestExamplePass(t *testing.T) {
	slog.Info("This will only be logged if the test fails2")
	slog.Info("This log will now appear since the test has failed2")
}
