package parser

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func LoadTest() string {

	dat, err := ioutil.ReadFile("../test-data/blame-test.txt")

	if err != nil {
		panic(err)
	}

	//	fmt.Println(string(dat))

	return string(dat)

}

func Test_Parse(t *testing.T) {

	lines := LoadTest()

	Parse(lines)

}

func Test_ParseHeader(t *testing.T) {

	file_lines := strings.Split(LoadTest(), "\n")

	lines := file_lines[:12]
	bl := BlameLines{lines, 0}
	header, bl := ParseHeader(bl)

	t.Log(header)

	assert.Equal(t, header.oid, "4e8a3451534e82b131a3c27fbcccadafb417de8f")
	assert.Equal(t, header.num_lines, 1)
	assert.Equal(t, header.author, "Todd Bush")

}

func Test_ParseLines(t *testing.T) {

	file_lines := strings.Split(LoadTest(), "\n")
	lines := file_lines[13:20]

	bl := BlameLines{lines, 0}

	extracted_lines, bl := ParseLines(bl, 1)

	assert.Equal(t, len(extracted_lines), 1)

}
