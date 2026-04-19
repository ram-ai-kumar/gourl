package internal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// OpenBrowser launches the provided URL in the default browser.
func OpenBrowser(url string) {
	// Check if we're in test mode and stub the opening
	if os.Getenv("GOURL_TEST_MODE") == "1" {
		fmt.Printf("🧪 TEST MODE: Would open %s (stubbed)\n", url)
		return
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "linux":
		// Check for WSL
		if _, err := exec.LookPath("wslview"); err == nil {
			cmd = exec.Command("wslview", url)
		} else {
			cmd = exec.Command("xdg-open", url)
		}
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("❌ Error opening URL: %v\n", err)
	} else {
		fmt.Printf("🚀 Opening %s\n", url)
	}
}
