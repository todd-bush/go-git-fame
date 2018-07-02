package git

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_git_list_files(t *testing.T) {

	files := git_list_files("master")

	assert.Equal(t, len(files) > 0, true)

	fmt.Printf("files: %v\n", files)

}

func Test_branch_exist(t *testing.T) {

	assert.Equal(t, branch_exists("master"), true)
	assert.Equal(t, branch_exists("no-branch-name"), false)
}
