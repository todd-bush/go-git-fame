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

func Test_GatherBlame(t *testing.T) {

	blame_results := GatherBlame("master")

	assert.NotNil(t, blame_results)

}

func Test_GatherCommits(t *testing.T) {
	test_result := GatherCommits()

	assert.NotNil(t, test_result)
	assert.Equal(t, len(test_result) > 1, true)

	for k, v := range test_result {

		log.Debugf("results include key=%s, val=%d", k, v)

		assert.NotNil(t, k)
		assert.NotNil(t, v)

	}
}
