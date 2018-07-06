package cmd

import (
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func Test_ExecuteProcessor(t *testing.T) {

	log.Debugln(" starting processor testing")

	os.Chdir("../")

	ExecuteProcessor()

}
