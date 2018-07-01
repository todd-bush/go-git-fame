package parser

import (
	//	"fmt"
	"github.com/matryer/is"
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

	is.New(t)

	lines := LoadTest()

	Parse(lines)

}

func Test_ParseHeader(t *testing.T) {
	is := is.New(t)

	file_lines := strings.Split(LoadTest(), "\n")

	lines := file_lines[:12]
	bl := BlameLines{lines, 0}
	header, bl := ParseHeader(bl)

	t.Log(header)

	is.Equal(header.oid, "4e8a3451534e82b131a3c27fbcccadafb417de8f")
	is.Equal(header.num_lines, 1)
	is.Equal(header.author, "Todd Bush")

}

func Test_ParseLines(t *testing.T) {
	is := is.New(t)

	file_lines := strings.Split(LoadTest(), "\n")
	lines := file_lines[13:20]

	bl := BlameLines{lines, 0}

	extracted_lines, bl := ParseLines(bl, 1)

	is.Equal(len(extracted_lines), 1)

}
