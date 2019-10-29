package git

import (
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

/*
	preforms a 'git ls-tree' on the branch arguments.
	returns a list of files.
*/
func GitListFiles(branch string) []string {

	var sb strings.Builder

	sb.WriteString("git ls-tree -r --name-only ")
	sb.WriteString(branch)

	r, e := executeGitCommand(sb.String())

	if e != nil {
		log.Fatalf("GitListFiles failed: %v", e)
	}

	return r

}

/*
Check to see if the branch argument exists
boolean if the branch exists or not
*/
func BranchExists(branch string) bool {
	var sb strings.Builder

	sb.WriteString("git show-ref ")
	sb.WriteString(branch)

	showOut, _ := exec.Command("bash", "-c", sb.String()).Output()

	return len(string(showOut)) > 0
}

/*
Performs a 'git blame' on the file argument
returns output as string slice
*/
func GitBlame(file string) []string {

	var sb strings.Builder

	cleanFile := strings.Replace(file, "\"", "", -1)

	escapeFile := cleanFile == file

	sb.WriteString("git blame -M -p -w -- ")
	if escapeFile {
		sb.WriteString("'")
	}
	sb.WriteString(cleanFile)

	if escapeFile {
		sb.WriteString("'")

	}

	log.Debugf("running Blame on file: %s", cleanFile)

	r, e := executeGitCommand(sb.String())

	if e != nil {
		return make([]string, 0)
	}

	return r
}

/*
Performs a 'git shortlog' on the current directory
returns output as string slice
*/
func GitShortLog() []string {
	shortCmd := "git log --pretty=short | git shortlog -nse"

	r, e := executeGitCommand(shortCmd)

	if e != nil {
		log.Fatalf("GitShortLog failed error: %v", e)
	}

	return r
}

func GitCurrentBranch() string {
	branchCmd := "git branch | grep \\* | cut -d ' ' -f2"

	results, _ := executeGitCommand(branchCmd)

	var result string

	if len(results) > 0 {
		result = results[0]
	}

	return result
}

func executeGitCommand(command string) ([]string, error) {

	gitOut, err := exec.Command("sh", "-c", command).Output()

	if err != nil {

		log.Errorf("Error while running GIT command \"%s\"\n %v\n", command, err)
		return nil, err
	}

	lines := strings.Split(string(gitOut), "\n")

	return lines, nil

}
