package git

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func git_list_files(branch string) []string {

	git_cmd := "git ls-tree -r --name-only " + branch

	ls_cmd := exec.Command("bash", "-c", git_cmd)

	stderr, _ := ls_cmd.StderrPipe()

	ls_out, err := ls_cmd.Output()

	if err != nil {
		slurp, _ := ioutil.ReadAll(stderr)
		fmt.Printf("slurp: %s\n", slurp)
		log.Fatalf("%T\n", err)
	}

	files := strings.Split(string(ls_out), "\n")

	return files

}
