package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_ExecuteProcessor(t *testing.T) {

	log.Debugln(" starting processor testing")

	os.Chdir("../")

	ExecuteProcessor()

	assert.Equal(t, true, true)

}

func Test_GatherCommits(t *testing.T) {
	test_result := GatherCommits()

	assert.NotNil(t, test_result)
	assert.Equal(t, len(test_result) > 1, true)
}
