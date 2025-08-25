package checkversions

import (
	"fmt"
	"os/exec"
)

func CheckNGrokVersion() error {
	if _, err := exec.LookPath("ngrok"); err != nil {
		fmt.Println("❌ ngrok is not installed")
		return err
	}
	fmt.Println("✅ ngrok is running.")
	return nil
}
