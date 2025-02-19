package client

import (
	"embed"
)

// Assets contains the embedded client assets
//
//go:embed src/*
var assets embed.FS
