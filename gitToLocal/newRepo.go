package gittolocal

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

	jsonrequests "github.com/hiabhi-cpu/cicdpipe/jsonRequests"
	"github.com/hiabhi-cpu/cicdpipe/mailing"
)

var logFile *os.File

const FILENAME = "localGitData.json"

func NewRepo() {

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
		logFile.WriteString(fmt.Sprintf("[%s] Failed to clone: %v\n", time.Now().Format(time.RFC3339), err))
		return
	}
	checkForLocalGitData(gitRepo)
	AddFiletoGitIgnore(gitRepo)
	// logFile.WriteString("Clonging repo" + string(out))
	logFile.WriteString(fmt.Sprintf("[%s] Clonging repo: %s\n", time.Now().Format(time.RFC3339), string(out)))
	fmt.Println("Cloning the repo to local")
	mailing.Mailing("New Webhook", "A new webhook has been created for "+gitRepo+" repository and cloned.", []string{})
}

func checkForLocalGitData(gitRepo string) {

	file, err := os.Create(FILENAME)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(jsonrequests.LocalJson{Git_Repo: gitRepo}); err != nil {
		fmt.Println("Error writing JSON:", err)
		return
	}

	fmt.Println("Data written to log.json successfully!")
}
