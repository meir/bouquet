package bouquet

import "path/filepath"

// Restore copies the backed up _core.asar back to the asarPath
func Restore(asarPath string) error {
	// move _core.asar to core.asar
	dirname := filepath.Dir(asarPath)
	backupPath := filepath.Join(dirname, "_core.asar")
	return copy(backupPath, asarPath)
}
