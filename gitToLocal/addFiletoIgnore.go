package gittolocal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AddFiletoGitIgnore(gitRepo string) error {
	fileName := ".gitignore"
	patternSplit := strings.Split(gitRepo, "/")
	pattern := patternSplit[len(patternSplit)-1]
	fmt.Println(pattern)
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Check if pattern already exists
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == pattern {
			fmt.Println("Pattern already exists in .gitignore")
			return nil
		}
	}

	// Append pattern at the end
	if _, err := file.WriteString("\n" + pattern + "\n"); err != nil {
		return err
	}
	fmt.Println("Added pattern to .gitignore:", pattern)
	return nil
}
