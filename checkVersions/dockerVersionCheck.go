package checkversions

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
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

	fmt.Println("✅ Docker is installed and running logging in")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dockerUser := os.Getenv("DOCKER_USER")
	dockerPass := os.Getenv("DOCKER_PAT")
	cmd = exec.Command("docker", "login", "-u", dockerUser, "-p", dockerPass)
	if err = cmd.Run(); err != nil {
		fmt.Println("Login to docker error")
		return err
	}
	fmt.Println("Docker Login Successful")
	return nil
}
