// Package app provides the top-level Wails application wiring, IPC bridge, and lifecycle management.
// The App struct is bound to the Wails frontend and its exported methods are callable from TypeScript.
package app

import "context"

// App is the Wails application struct. It is constructed once in main and bound to the Wails runtime.
// Exported methods on App are automatically exposed to the frontend as TypeScript async functions
// by the Wails code-generation tool.
type App struct {
	ctx context.Context
}

// New constructs a new App instance ready to be bound to Wails.
func New() *App {
	return &App{}
}

// OnStartup is called by the Wails runtime immediately after the application window is created.
// The provided context carries the Wails runtime context and is used for the lifetime of the application.
func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx
}
