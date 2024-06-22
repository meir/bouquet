package bouquet

import (
	"errors"
	"strings"

	version "github.com/meir/bouquet"
	"github.com/meir/bouquet/pkg/asar"
)

// Version will return the installed version of bouquet and the current version of the binary
func Version(asarPath string) (string, string, error) {
	a, err := asar.NewAsar(asarPath)
	if err != nil {
		return "", "", err
	}

	if a == nil {
		return "", "", errors.New("header was nil")
	}

	installedVersion := a.Header.Get("bouquet/VERSION")
	version := strings.TrimSpace(version.VERSION)
	if installedVersion == nil {
		return "", version, nil
	}

	installed := strings.TrimSpace(string(installedVersion.Content()))

	return installed, version, nil
}
