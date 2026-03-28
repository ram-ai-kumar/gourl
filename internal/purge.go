package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

// PurgeApp removes the running binary from the system.
func PurgeApp() {
	// 1. Locate the executable
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("❌ Error: Could not locate gourl executable: %v\n", err)
		os.Exit(1)
	}

	// 2. Resolve symlinks
	realPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		realPath = exePath // Fallback to raw path if symlink resolution fails
	}

	// 3. Remove the binary
	err = os.Remove(realPath)
	if err != nil {
		fmt.Printf("❌ Error: Failed to remove binary: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ gourl has been successfully uninstalled.")
	fmt.Println("💡 Note: Your project-specific .cache/ folders were preserved.")
}
