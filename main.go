package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	checkversions "github.com/hiabhi-cpu/cicdpipe/checkVersions"
	gittolocal "github.com/hiabhi-cpu/cicdpipe/gitToLocal"
	gitlib "github.com/hiabhi-cpu/gitwebhook/gitLib"
	"github.com/joho/godotenv"
)

var logFile *os.File
var ngrokCmd *exec.Cmd

func main() {
	if checkversions.CheckGitVersion() != nil {
		panic("No git in the local system install git first")
	}
	if checkversions.CheckDockerVersion() != nil {
		panic("Install docker and check it's working")
	}
	if checkversions.CheckNGrokVersion() != nil {
		panic("Install docker and check it's working")
	}
	fmt.Println("Hello")
	var err error
	logFile, err = os.OpenFile("mainLogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer logFile.Close()

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	user_PAT := os.Getenv("GIT_PAT")
	gitRepo := os.Getenv("GIT_REPO")
	rev_url := os.Getenv("REV_URL")
	err = gitlib.GetOrCreateWebhook(gitRepo, user_PAT, rev_url)

	if err != nil {
		fmt.Println("Failed to Get or Create Webhook", err)
		return
	}

	if err = runGrokCommand(); err != nil {
		fmt.Println(err)
		return
	}

	go handleShutdown()

	http.HandleFunc("POST /webhook", webHookHandler)
	http.HandleFunc("GET /webhook", getWebhookHandler)
	fmt.Println("Listening to port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("After listener")
}

func handleShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c // Wait for Ctrl+C
	logFile.WriteString(fmt.Sprintf("[%s] Received shutdown signal. Killing ngrok...\n", time.Now().Format(time.RFC3339)))

	if ngrokCmd != nil && ngrokCmd.Process != nil {
		ngrokCmd.Process.Kill()
		logFile.WriteString(fmt.Sprintf("[%s] ngrok process killed\n", time.Now().Format(time.RFC3339)))
	}
	os.Exit(0)
}

func webHookHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	logFile.WriteString(fmt.Sprintf("[%s] Received POST request\n", time.Now().Format(time.RFC3339)))
	fmt.Println("Received POST request")
	fmt.Println(string(body))

	go gittolocal.GetGitToLocal(body)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook created"))
}

func getWebhookHandler(w http.ResponseWriter, r *http.Request) {
	logFile.WriteString(fmt.Sprintf("[%s] Received GET request\n", time.Now().Format(time.RFC3339)))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GET Webhook created"))
}

func runGrokCommand() error {
	// Start ngrok
	ngrokCmd = exec.Command("ngrok", "http", "--url=raccoon-model-reasonably.ngrok-free.app", "8080") // <-- fixed wrong flag
	ngrokCmd.Stdout = logFile
	ngrokCmd.Stderr = logFile
	if err := ngrokCmd.Start(); err != nil {
		logFile.WriteString(fmt.Sprintf("[%s] Failed to start ngrok: %v\n", time.Now().Format(time.RFC3339), err))
		return err
	}
	logFile.WriteString(fmt.Sprintf("[%s] ngrok started with PID: %d\n", time.Now().Format(time.RFC3339), ngrokCmd.Process.Pid))
	return nil
}
