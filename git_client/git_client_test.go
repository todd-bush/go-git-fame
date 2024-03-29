package git

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GitListFiles(t *testing.T) {

	files := GitListFiles("master")

	assert.Equal(t, len(files) > 0, true)

	fmt.Printf("files: %v\n", files)

}

func Test_BranchExists(t *testing.T) {

	assert.Equal(t, BranchExists("master"), true)
	assert.Equal(t, BranchExists("no-branch-name"), false)
}

func Test_GitBlame(t *testing.T) {
	lines := GitBlame("../Makefile") // include path

	assert.NotNil(t, lines)
	assert.Equal(t, len(lines) > 10, true)

}

func Test_GitShortLog(t *testing.T) {

	lines := GitShortLog()

	//t.Log(lines)

	assert.NotNil(t, lines)
	assert.Equal(t, len(lines) > 1, true)
}

func Test_GitCurrentBranch(t *testing.T) {
	result := GitCurrentBranch()

	assert.NotNil(t, result)
	t.Log(result)
}

func Test_GitCommitDate(t *testing.T) {
	result := GitCommitDate()

	assert.NotNil(t, result)
	assert.Equal(t, len(result) > 1, true)
	t.Log(result)

}
