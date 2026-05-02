//go:build desktop

// Package main is the entry point for the Zevaro daemon and GUI.
package main

import "embed"

// assets holds the compiled frontend bundle, embedded at build time.
// The frontend/dist directory is populated by `pnpm run build` and copied here
// by the Makefile `build` target before Go compilation.
//
//go:embed all:frontend/dist
var assets embed.FS
