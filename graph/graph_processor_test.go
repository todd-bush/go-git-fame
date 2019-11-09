package graph

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CollectCommits(t *testing.T) {

	//logrus.SetLevel(logrus.InfoLevel)

	commits := CollectCommits()

	assert.NotNil(t, commits)

	var lastDate time.Time

	for i, c := range commits {

		fmt.Printf("%d - %+v\n", i, c)

		if i != 0 {
			assert.Equal(t, lastDate.After(c.date), true)
		}

		lastDate = c.date

	}

}

func Test_GraphCommits(t *testing.T) {
	commits := CollectCommits()

	assert.NotNil(t, commits)

	GraphCommits(commits)

	_, fErr := os.Stat(CommitByDateFileName)
	assert.Nil(t, fErr)

	dErr := os.Remove(CommitByDateFileName)
	if dErr != nil {
		t.Log(dErr)
	}

}
