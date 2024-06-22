package bouquet

import (
	"os"
	"path/filepath"
)

// Backup backs up the core.asar file to _core.asar
func Backup(asarPath string) error {
	// move core.asar to _core.asar if it doesnt exists yet
	dirname := filepath.Dir(asarPath)
	backupPath := filepath.Join(dirname, "_core.asar")

	if _, err := os.Stat(backupPath); err == nil {
		return nil
	}

	return copy(asarPath, backupPath)
}
