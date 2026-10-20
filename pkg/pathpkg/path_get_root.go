package pathpkg

import (
	"path/filepath"
	"runtime"
)

// PathGetRoot just use in root/internal/cmd/main.go RETURN ./root
func PathGetRoot() string {
	// Get the path of the current file
	_, currentFile, _, _ := runtime.Caller(0)
	// Move up to the root directory (adjust the number of '../' based on your project structure)
	projectRoot := filepath.Join(filepath.Dir(currentFile), "../..")
	// Clean the resulting path
	return filepath.Clean(projectRoot)
}
