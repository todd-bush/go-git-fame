package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_Main(t *testing.T) {
	log.SetLevel(log.WarnLevel)
}

func Test_ExecuteProcessor(t *testing.T) {

	log.Debugln(" starting processor testing")

	os.Chdir("../")

	result := ExecuteProcessor("master")

	assert.Equal(t, len(result) > 0, true)

}

func Test_GatherBlame(t *testing.T) {

	blame_results := GatherBlame("master")

	assert.NotNil(t, blame_results)
	assert.Equal(t, len(blame_results) > 0, true)

	for _, result := range blame_results {
		log.Debugf("blame results for file %s", result.file)
	}

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
