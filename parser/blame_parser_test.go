package parser

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	log.SetLevel(log.ErrorLevel)
	log.SetOutput(os.Stdout)
	code := m.Run()
	os.Exit(code)
}

func LoadTest() string {

	dat, err := ioutil.ReadFile("../test-data/blame-test.txt")

	if err != nil {
		panic(err)
	}

	return string(dat)

}

func Test_Parse(t *testing.T) {

	assert := assert.New(t)
	allLines := LoadTest()

	lines := strings.Split(allLines, "\n")

	data := Parse(lines)

	//t.Log(fmt.Sprintf("Parse Data: %v\n", data))

	assert.Equal(27, len(data))
	assert.Equal(data[0].oid, "4e8a3451534e82b131a3c27fbcccadafb417de8f")

}

func Test_ParseHeader(t *testing.T) {

	fileLines := strings.Split(LoadTest(), "\n")

	lines := fileLines[:12]
	bl := BlameLines{lines, 0}
	header, bl := ParseHeader(bl)

	t.Log(header)

	assert.Equal(t, header.oid, "4e8a3451534e82b131a3c27fbcccadafb417de8f")
	assert.Equal(t, header.NumLines, 1)
	assert.Equal(t, header.Author, "Todd Bush")

}

func Test_ParseLines(t *testing.T) {

	fileLines := strings.Split(LoadTest(), "\n")
	lines := fileLines[13:20]

	bl := BlameLines{lines, 0}

	bl = ParseLines(bl, 1)

	assert.Equal(t, bl.indexPtr, 1)

}
