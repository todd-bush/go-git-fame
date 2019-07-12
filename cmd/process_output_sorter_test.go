package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortCommits(t *testing.T) {
	testSlice := make([]ProcessOutput, 3)
	testSlice[0] = ProcessOutput{author: "first", loc: 100, commits: 100, fileCount: 100}
	testSlice[2] = ProcessOutput{author: "second", loc: 200, commits: 200, fileCount: 200}
	testSlice[1] = ProcessOutput{author: "third", loc: 300, commits: 300, fileCount: 300}

	By(ByCommits).Sort(testSlice)

	assert.Equal(t, testSlice[0].commits, 300)
	assert.Equal(t, testSlice[1].commits, 200)
	assert.Equal(t, testSlice[2].commits, 100)
}

func TestSortLoc(t *testing.T) {

	testSlice := make([]ProcessOutput, 3)
	testSlice[0] = ProcessOutput{author: "first", loc: 100, commits: 100, fileCount: 100}
	testSlice[2] = ProcessOutput{author: "second", loc: 200, commits: 200, fileCount: 200}
	testSlice[1] = ProcessOutput{author: "third", loc: 300, commits: 300, fileCount: 300}

	By(ByLoc).Sort(testSlice)

	assert.Equal(t, testSlice[0].loc, 300)
	assert.Equal(t, testSlice[1].loc, 200)
	assert.Equal(t, testSlice[2].loc, 100)

}

func TestSortFiles(t *testing.T) {

	testSlice := make([]ProcessOutput, 3)
	testSlice[0] = ProcessOutput{author: "first", loc: 100, commits: 100, fileCount: 100}
	testSlice[2] = ProcessOutput{author: "second", loc: 200, commits: 200, fileCount: 200}
	testSlice[1] = ProcessOutput{author: "third", loc: 300, commits: 300, fileCount: 300}

	By(ByFiles).Sort(testSlice)

	assert.Equal(t, testSlice[0].fileCount, 300)
	assert.Equal(t, testSlice[1].fileCount, 200)
	assert.Equal(t, testSlice[2].fileCount, 100)

}
