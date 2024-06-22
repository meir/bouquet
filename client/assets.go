package client

import (
	"embed"
)

// Assets contains the embedded client assets
//
//go:embed src/* patches/*
var assets embed.FS
