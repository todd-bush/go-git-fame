package git

import (
	"fmt"
	"testing"
)

func Test_git_list_files(t *testing.T) {
	files := git_list_files("master")

	if len(files) <= 0 {
		t.Fatal("Expecting files to be larger than zero")
	} else {
		fmt.Printf("files: %v\n", files)
	}

}
