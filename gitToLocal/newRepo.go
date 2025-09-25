package gittolocal

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/hiabhi-cpu/cicdpipe/dockerbuild"
	jsonrequests "github.com/hiabhi-cpu/cicdpipe/jsonRequests"
	"github.com/hiabhi-cpu/cicdpipe/mailing"
)

var logFile *os.File
var gitRepo string
var oldRepo string

const FILENAME = "localGitData.json"

func NewRepo() {

	gitRepo = os.Getenv("GIT_REPO")

	logFile, err := os.OpenFile("mainLogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer logFile.Close()

	checkForLocalGitData(gitRepo)
	AddFiletoGitIgnore(gitRepo)

	out, err := exec.Command("git", "clone", "https://"+gitRepo).Output()
	if err != nil {
		fmt.Println("Error in Cloning")
		logFile.WriteString(fmt.Sprintf("[%s] Failed to clone: %v\n", time.Now().Format(time.RFC3339), err))
		return
	}

	// logFile.WriteString("Clonging repo" + string(out))
	logFile.WriteString(fmt.Sprintf("[%s] Clonging repo: %s\n", time.Now().Format(time.RFC3339), string(out)))
	fmt.Println("Cloning the repo to local")
	dockerbuild.CheckDockerInRepo()
	mailing.Mailing("New Webhook", "A new webhook has been created for "+gitRepo+" repository and cloned.", []string{})
}

func checkForLocalGitData(gitRepo string) {

	var existing jsonrequests.LocalJson

	// Step 1: Check if file exists
	if _, err := os.Stat(FILENAME); err == nil {
		// File exists → read and unmarshal
		data, err := os.ReadFile(FILENAME)
		if err == nil && len(data) > 0 {
			_ = json.Unmarshal(data, &existing)
		}
	}

	// Step 2: Compare old vs new
	if existing.Git_Repo == gitRepo {
		fmt.Println("No changes, skipping write.")
		return
	}

	// Step 3: Write new data (overwrite file)
	file, err := os.Create(FILENAME)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	oldRepo = getFileName(existing.Git_Repo)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(jsonrequests.LocalJson{Git_Repo: gitRepo}); err != nil {
		fmt.Println("Error writing JSON:", err)
		return
	}

	fmt.Println("JSON data updated in", FILENAME)
	removeOldRepo(oldRepo)
	RemoveOldRepoFromGitIgnore(oldRepo)
}

func removeOldRepo(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("Directory does not exist:", dir)
		return
	}

	if err := os.RemoveAll(dir); err != nil {
		fmt.Println("Error removing directory:", err)
		return
	}

	fmt.Println("Directory", dir, "removed successfully")
}
func getFileName(existing string) string {

	str := strings.Split(existing, "/")
	return str[len(str)-1]
}
