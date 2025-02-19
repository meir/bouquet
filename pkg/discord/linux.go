//go:build linux

package discord

import (
	"os"
	"path/filepath"
)

// GetPath gets the location of the Discord app
func GetPath() (string, error) {
	f := filepath.Join(os.Getenv("HOME"), "/.config/discord/app.asar")

	// check if asar file exists
	_, err := os.Stat(f)
	switch err {
	case nil:
		return f, nil
	case os.ErrNotExist:
		// TODO: check if app is installed somewhere else
		fallthrough
	default:
		return "", nil
	}
}
