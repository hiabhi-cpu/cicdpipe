package gittolocal

import (
	"fmt"
	"os"
	"os/exec"
)

var logFile *os.File

const FILENAME = "localGitData.json"

func NewRepo() {
	checkForLocalGitData()
	gitRepo := os.Getenv("GIT_REPO")

	logFile, err := os.OpenFile("mainLogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer logFile.Close()

	out, err := exec.Command("git", "clone", "https://"+gitRepo).Output()
	if err != nil {
		fmt.Println("Error in Cloning")
		return
	}
	AddFiletoGitIgnore(gitRepo)
	logFile.WriteString("Clonging repo" + string(out))
	fmt.Println("Cloning the repo to local")

}

func checkForLocalGitData() {
	if _, err := os.Stat(FILENAME); os.IsNotExist(err) {
		fmt.Println("File does not exist, creating...")
		f, _ := os.Create(FILENAME)
		defer f.Close()
	} else {
		fmt.Println("File already exists")
	}
}
