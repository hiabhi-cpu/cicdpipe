package gittolocal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/hiabhi-cpu/cicdpipe/mailing"
)

func NewCommit() {
	gitRepo := os.Getenv("GIT_REPO")

	logFile, err := os.OpenFile("mainLogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer logFile.Close()
	patternSplit := strings.Split(gitRepo, "/")
	folderName := patternSplit[len(patternSplit)-1]

	cmd := exec.Command("git", "pull")
	cmd.Dir = "./" + folderName
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error in Pulling")
		logFile.WriteString(fmt.Sprintf("[%s] Failed to pull: %s\n", time.Now().Format(time.RFC3339), err))
		return
	}
	AddFiletoGitIgnore(gitRepo)
	// logFile.WriteString("Pulling repo" + string(out))
	logFile.WriteString(fmt.Sprintf("[%s] Pulling repo: %v\n", time.Now().Format(time.RFC3339), string(out)))
	fmt.Println("Pulling new commit repo to local")
	mailing.Mailing("New Commit", "A new commit has been created for "+gitRepo+" repository and pulled.", []string{})
}
