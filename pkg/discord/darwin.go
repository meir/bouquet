//go:build darwin

package discord

import (
	"os"
)

// GetPath gets the location of the Discord app
func GetPath() (string, error) {
	f := "/Applications/Discord.app/Contents/Resources/app.asar"

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
