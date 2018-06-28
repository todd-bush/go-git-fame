package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func git_list_files(dir string, branch string) []string {

	git_cmd := "git -C " + dir + " ls-files --with-tree " + branch

	ls_cmd := exec.Command("bash", "-c", git_cmd)

	stderr, _ := ls_cmd.StderrPipe()

	ls_out, err := ls_cmd.Output()

	if err != nil {
		slurp, _ := ioutil.ReadAll(stderr)
		fmt.Printf("slurp:  %s\n", slurp)
		log.Fatalf("%T\n", err)
	}

	files := strings.Split(string(ls_out), "\n")

	return files
}

func main() {

	// get the command line args
	args := os.Args[1:]
	fmt.Println(args)
}
