package git

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
)

func handle_error(err error, stderr io.Reader) {

	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("slurp: %s\n", slurp)
	log.Fatalf("%T\n", err)

}

func ExecuteGitCommand(command string) []string {

	git_cmd := exec.Command("bash", "-c", command)

	stderr, _ := git_cmd.StderrPipe()

	git_out, err := git_cmd.Output()

	if err != nil {
		handle_error(err, stderr)
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

	show_cmd := exec.Command("bash", "-c", git_cmd)

	show_out, _ := show_cmd.Output()

	return len(string(show_out)) > 0
}

/*
Performs a 'git blame' on the file argument
returns output as string slice
*/
func GitBlame(file string) []string {
	git_cmd := fmt.Sprintf("git blame -M -p -w -- '%s'", file)

	return ExecuteGitCommand(git_cmd)
}

/*
Performs a 'git shortlog' on the current directory
returns output as string slice
*/
func GitShortLog() []string {
	git_cmd := "git shortlog -s -e"

	return ExecuteGitCommand(git_cmd)
}
