package git

import (
	"fmt"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

/*
	preforms a 'git ls-tree' on the branch arguments.
	returns a list of files.
*/
func GitListFiles(branch string) []string {

	gitCmd := fmt.Sprintf("git ls-tree -r --name-only %s", branch)

	return executeGitCommand(gitCmd)

}

/*
Check to see if the branch argument exists
boolean if the branch exists or not
*/
func BranchExists(branch string) bool {
	gitCmd := fmt.Sprintf("git show-ref %s", branch)

	showOut, _ := exec.Command("bash", "-c", gitCmd).Output()

	return len(string(showOut)) > 0
}

/*
Performs a 'git blame' on the file argument
returns output as string slice
*/
func GitBlame(file string) []string {
	gitCmd := fmt.Sprintf("git blame -M -p -w -- '%s'", file)

	log.Debugf("running Blame on file: %s", file)

	return executeGitCommand(gitCmd)
}

/*
Performs a 'git shortlog' on the current directory
returns output as string slice
*/
func GitShortLog() []string {
	shortCmd := "git log --pretty=short | git shortlog -nse"

	return executeGitCommand(shortCmd)
}

func GitCurrentBranch() string {
	branchCmd := "git branch | grep \\* | cut -d ' ' -f2"

	results := executeGitCommand(branchCmd)

	var result string

	if len(results) > 0 {
		result = results[0]
	}

	return result
}

func executeGitCommand(command string) []string {

	gitOut, err := exec.Command("sh", "-c", command).Output()

	if err != nil {
		log.Fatalf("%T\n", err)
	}

	lines := strings.Split(string(gitOut), "\n")

	return lines

}
