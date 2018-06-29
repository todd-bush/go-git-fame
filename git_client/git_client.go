package git

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func handle_error(err error, stderr io.Reader) {

	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("slurp: %s\n", slurp)
	log.Fatalf("%T\n", err)

}

func git_list_files(branch string) []string {

	git_cmd := "git ls-tree -r --name-only " + branch

	ls_cmd := exec.Command("bash", "-c", git_cmd)

	stderr, _ := ls_cmd.StderrPipe()

	ls_out, err := ls_cmd.Output()

	if err != nil {
		handle_error(err, stderr)
	}

	files := strings.Split(string(ls_out), "\n")

	return files

}

func branch_exists(branch string) bool {
	git_cmd := "git show-ref " + branch

	show_cmd := exec.Command("bash", "-c", git_cmd)

	show_out, _ := show_cmd.Output()

	//fmt.Println(string(show_out))

	return len(string(show_out)) > 0
}
