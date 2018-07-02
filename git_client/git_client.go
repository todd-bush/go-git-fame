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

/*
	preforms a 'git ls-tree' on the branch arguments.
	returns a list of files.
*/
func GitListFiles(branch string) []string {

	git_cmd := fmt.Sprintf("git ls-tree -r --name-only %s", branch)

	ls_cmd := exec.Command("bash", "-c", git_cmd)

	stderr, _ := ls_cmd.StderrPipe()

	ls_out, err := ls_cmd.Output()

	if err != nil {
		handle_error(err, stderr)
	}

	files := strings.Split(string(ls_out), "\n")

	return files

}

/*
Check to see if the branch argument exists
boolean if the branch exists or not
*/
func BranchExists(branch string) bool {
	git_cmd := fmt.Sprintf("git show-ref %s", branch)

	show_cmd := exec.Command("bash", "-c", git_cmd)

	show_out, _ := show_cmd.Output()

	log.Debug(string(show_out))

	return len(string(show_out)) > 0
}

/*
Performs a 'git blame' on the file argument
returns the blame output
*/
func GitBlame(file string) []string {
	git_cmd := fmt.Sprintf("git blame -M -p -w -- '%s'", file)

	fmt.Println(git_cmd)

	blame_cmd := exec.Command("bash", "-c", git_cmd)

	stderr, _ := blame_cmd.StderrPipe()

	blame_out, err := blame_cmd.Output()

	if err != nil {
		handle_error(err, stderr)
	}

	lines := strings.Split(string(blame_out), "\n")

	return lines
}
