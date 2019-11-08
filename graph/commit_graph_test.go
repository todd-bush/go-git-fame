package graph

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GraphCommitsByTime(t *testing.T) {
	input := make([]CommitsByDate, 4)
	input[0] = CommitsByDate{time.Now(), float64(15)}
	input[1] = CommitsByDate{time.Now().Add(-time.Hour * 24), float64(14)}
	input[2] = CommitsByDate{time.Now().Add(-time.Hour * 48), float64(13)}
	input[3] = CommitsByDate{time.Now().Add(-time.Hour * 72), float64(12)}

	err := GraphCommitsByTime(input)

	assert.Nil(t, err)

	_, fErr := os.Stat(CommitByDateFileName)
	assert.Nil(t, fErr)

	dErr := os.Remove(CommitByDateFileName)
	if dErr != nil {
		t.Log(dErr)
	}

}
