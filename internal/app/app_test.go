// Package app tests the Wails application lifecycle wiring.
package app

import (
	"context"
	"testing"
)

// TestNew verifies that New returns a non-nil App struct.
func TestNew(t *testing.T) {
	t.Parallel()

	a := New()
	if a == nil {
		t.Fatal("New() returned nil; expected a non-nil *App")
	}
}

// TestOnStartup verifies that OnStartup stores the context without panicking.
func TestOnStartup(t *testing.T) {
	t.Parallel()

	a := New()
	ctx := context.Background()
	a.OnStartup(ctx) // must not panic

	if a.ctx != ctx {
		t.Error("OnStartup did not store the provided context")
	}
}
