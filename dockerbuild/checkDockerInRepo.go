package dockerbuild

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/hiabhi-cpu/cicdpipe/mailing"
)

func CheckDockerInRepo() {
	logFile, err := os.OpenFile("mainLogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	gitRepo := os.Getenv("GIT_REPO")
	dockerVersion := os.Getenv("DOCKER_VERSION")
	fileName := getFileName(gitRepo)
	if !fileExists(fileName + "/docker-compose.yml") {
		fmt.Println("Dokcer-compose not present")
		if !fileExists(fileName + "/dockerfile") {
			fmt.Println("Dockerfile also not present")
			mailing.Mailing("Docker Not found", "Neither docker-compose nor dockerfile is present in the repo please add one", []string{})
			logFile.WriteString(fmt.Sprintf("[%s] dockerfile or docker-compose not found\n", time.Now().Format(time.RFC3339)))
			return
		} else {
			fmt.Println("Dockerfile is present")
			logFile.WriteString(fmt.Sprintf("[%s] dockerfile found\n", time.Now().Format(time.RFC3339)))
			dockerfileRun(fileName, dockerVersion)
			return
		}
	} else {
		fmt.Println("Docker compose present")
		logFile.WriteString(fmt.Sprintf("[%s] docker-compose found\n", time.Now().Format(time.RFC3339)))
		dockercomposeRun(fileName, dockerVersion)

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

func dockerfileRun(fileName, dockerVersion string) {
	fmt.Println("Building docker file ")
	fmt.Println("Building docker file ")
	cmd := exec.Command("docker", "build", "-t", fileName+":"+dockerVersion, ".")
	cmd.Dir = "./" + fileName
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error in building dockerfile")
		return
	}
	fmt.Println(string(out))
}

func dockercomposeRun(fileName, dockerVersion string) {
	fmt.Println("Building docker compose")
}
