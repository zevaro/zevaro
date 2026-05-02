//go:build !desktop

// Package main is the entry point for the Zevaro daemon and GUI.
package main

import "embed"

// assets is an empty embed.FS in development and test modes.
// In development, Wails serves assets via the Vite dev server proxy.
// In production, build with -tags desktop to embed frontend/dist.
var assets embed.FS
