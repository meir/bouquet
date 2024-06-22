//go:build linux

package discord

import (
	"os"
	"path/filepath"
)

// GetPath gets the location of the Discord app
func GetPath() (string, error) {
	path := filepath.Join(os.Getenv("HOME"), "/.config/discord")
	folders, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	for _, folder := range folders {
		if folder.IsDir() && version_regex.MatchString(folder.Name()) {
			return filepath.Join(path, folder.Name()), nil
		}
	}

	return "", nil
}

// GetASARPath gets the path to the ASAR file within the Discord app
func GetASARPath(path string) string {
	return path + "/modules/discord_desktop_core/core.asar"
}
