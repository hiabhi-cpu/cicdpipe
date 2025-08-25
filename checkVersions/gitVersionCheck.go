package checkversions

import (
	"fmt"
	"os/exec"
	"strings"
)

func CheckGitVersion() error {
	out, err := exec.Command("git", "--version").Output()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	version := strings.TrimSpace(string(out))
	fmt.Println("✅ Git version:", version)
	return nil
}
