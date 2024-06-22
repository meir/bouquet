//go:build windows

package discord

// GetPath gets the location of the Discord app
func GetPath() (string, error) {
	return "", nil
}

// GetASARPath gets the path to the ASAR file within the Discord app
func GetASARPath(path string) string {
	return ""
}
