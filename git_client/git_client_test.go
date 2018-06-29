package git

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func Test_git_list_files(t *testing.T) {

	is := is.New(t)

	files := git_list_files("master")

	is.True(len(files) > 0)

	fmt.Printf("files: %v\n", files)

}
