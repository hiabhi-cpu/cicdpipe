package gittolocal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const fileName = ".gitignore"

func AddFiletoGitIgnore(gitRepo string) error {

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

func RemoveOldRepoFromGitIgnore(gitrepo string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Keep all lines except the one that matches target
		if strings.TrimSpace(line) != gitrepo && len(strings.TrimSpace(line)) != 0 {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// Rewrite the file without the target line
	return os.WriteFile(fileName, []byte(strings.Join(lines, "\n")+"\n"), 0644)
}
