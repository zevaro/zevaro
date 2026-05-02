// Package main is the entry point for the Zevaro daemon and GUI.
// It initializes structured logging, constructs the Wails application, and starts the event loop.
package main

import (
	"log/slog"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/zevaro/zevaro/internal/app"
)

// defaultLogLevel is the log level used until configuration is loaded from config.yaml.
// Future tasks will replace this placeholder with a value read from the config layer.
const defaultLogLevel = slog.LevelInfo

// newLogger constructs the application-wide structured logger.
// In production (when ZEVARO_ENV=production), it uses a JSON handler to stderr.
// In all other modes it uses a human-readable text handler.
// The logger is set as the default slog logger for the process.
func newLogger() *slog.Logger {
	var handler slog.Handler
	if os.Getenv("ZEVARO_ENV") == "production" {
		handler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			Level: defaultLogLevel,
		})
	} else {
		handler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: defaultLogLevel,
		})
	}
	return slog.New(handler)
}

func main() {
	logger := newLogger()
	slog.SetDefault(logger)

	zevaroApp := app.New()

	if err := wails.Run(&options.App{
		Title:            "Zevaro",
		Width:            1280,
		Height:           800,
		MinWidth:         1024,
		MinHeight:        680,
		BackgroundColour: &options.RGBA{R: 11, G: 13, B: 16, A: 255},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: zevaroApp.OnStartup,
		Bind: []interface{}{
			zevaroApp,
		},
	}); err != nil {
		logger.Error("application exited with error", "error", err)
		os.Exit(1)
	}
}
