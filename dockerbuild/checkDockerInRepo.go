package dockerbuild

import (
	"fmt"
	"os"
	"strings"
)

func CheckDockerInRepo() {
	gitRepo := os.Getenv("GIT_REPO")
	fileName := getFileName(gitRepo)
	if !fileExists(fileName + "/docker-compose.yml") {
		fmt.Println("Dokcer-compose not present")
		if !fileExists(fileName + "/dockerfile") {
			fmt.Println("Dockerfile also not present")
		} else {
			fmt.Println("Dockerfile is present")
		}
	} else {
		fmt.Println("Docker compose present")

	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
func getFileName(existing string) string {

	str := strings.Split(existing, "/")
	return str[len(str)-1]
}
