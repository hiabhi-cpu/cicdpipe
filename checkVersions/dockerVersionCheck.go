package checkversions

import (
	"fmt"
	"os/exec"
)

func CheckDockerVersion() error {
	if _, err := exec.LookPath("docker"); err != nil {
		fmt.Println("❌ Docker is not installed")
		return err
	}

	// Check if Docker daemon is running
	cmd := exec.Command("docker", "info")
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Docker daemon is not running:", err)
		return err
	}

	fmt.Println("✅ Docker is installed and running")
	return nil
}
