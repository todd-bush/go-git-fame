package git

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

func ExecuteGitCommand(command string) []string {

	git_out, err := exec.Command("sh", "-c", command).Output()

	if err != nil {
		log.Fatalf("%T\n", err)
	}

	lines := strings.Split(string(git_out), "\n")

	return lines

}

/*
	preforms a 'git ls-tree' on the branch arguments.
	returns a list of files.
*/
func GitListFiles(branch string) []string {

	git_cmd := fmt.Sprintf("git ls-tree -r --name-only %s", branch)

	return ExecuteGitCommand(git_cmd)

}

/*
Check to see if the branch argument exists
boolean if the branch exists or not
*/
func BranchExists(branch string) bool {
	git_cmd := fmt.Sprintf("git show-ref %s", branch)

	show_out, _ := exec.Command("bash", "-c", git_cmd).Output()

	return len(string(show_out)) > 0
}

/*
Performs a 'git blame' on the file argument
returns output as string slice
*/
func GitBlame(file string) []string {
	git_cmd := fmt.Sprintf("git blame -M -p -w -- '%s'", file)

	log.Debugf("running Blame on file: %s", file)

	return ExecuteGitCommand(git_cmd)
}

/*
Performs a 'git shortlog' on the current directory
returns output as string slice
*/
func GitShortLog() []string {
	short_cmd := "git log --pretty=short | git shortlog -nse"

	return ExecuteGitCommand(short_cmd)
}
