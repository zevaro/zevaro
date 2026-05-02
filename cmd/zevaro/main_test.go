// Package main tests the Zevaro entry-point wiring.
package main

import (
	"testing"

	"github.com/zevaro/zevaro/internal/app"
)

// TestAppConstruction verifies that the Wails App struct can be constructed without error.
// It does not launch the GUI; it only exercises the wiring that runs before wails.Run.
func TestAppConstruction(t *testing.T) {
	t.Parallel()

	a := app.New()
	if a == nil {
		t.Fatal("app.New() returned nil; expected a non-nil *app.App")
	}
}

// TestNewLogger verifies that newLogger returns a non-nil logger in text mode (default).
func TestNewLogger(t *testing.T) {
	t.Parallel()

	logger := newLogger()
	if logger == nil {
		t.Fatal("newLogger() returned nil; expected a non-nil *slog.Logger")
	}
}

// TestNewLoggerProductionMode verifies that newLogger returns a non-nil JSON logger
// when ZEVARO_ENV=production.
func TestNewLoggerProductionMode(t *testing.T) {
	// Cannot run in parallel because it mutates the environment.
	t.Setenv("ZEVARO_ENV", "production")

	logger := newLogger()
	if logger == nil {
		t.Fatal("newLogger() returned nil in production mode; expected a non-nil *slog.Logger")
	}
}
