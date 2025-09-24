package gittolocal

import (
	"encoding/json"
	"fmt"

	jsonrequests "github.com/hiabhi-cpu/cicdpipe/jsonRequests"
)

func GetGitToLocal(body []byte) {

	var newWebhook jsonrequests.NewHookJson
	var newCommit jsonrequests.NewCommitJson

	if err := json.Unmarshal(body, &newWebhook); err != nil {
		fmt.Println("No new webhook created")
	}

	if err := json.Unmarshal(body, &newCommit); err != nil {
		fmt.Println("No new commit")
	}

	if newWebhook.Zen != "" {
		NewRepo()
		// fmt.Println(newWebhook.Zen)
	}
	if newCommit.Ref != "" {
		NewCommit()
		// fmt.Println(newCommit)
	}
}
