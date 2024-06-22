package discord

import (
	"path/filepath"
	"regexp"
)

var version_regex = regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+`)

func GetVersion() (string, error) {
	path, err := GetPath()
	if err != nil {
		return "", err
	}
	return filepath.Base(path), nil
}
