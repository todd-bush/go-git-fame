package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type BlameData struct {
	oid         string
	author      string
	num_lines   int
	mail        string
	time        string
	tz          string
	other_lines []string
}

func Parse(all_lines string) {

	lines := strings.Split(all_lines, "\n")

	//chunks := make(map[string]string)

	for {
		header := ParseHeader(lines)
		fmt.Println(header)
		fmt.Println(len(lines))
		break
	}

}

func ParseHeader(lines []string) BlameData {
	r, _ := regexp.Compile(`(?m)^([0-9a-f]{40}) (\d+) (\d+) (\d+)$`)

	headerline := splice(lines)

	pieces := r.FindStringSubmatch(headerline)
	numlines, _ := strconv.Atoi(pieces[4])

	bd := BlameData{oid: pieces[1], num_lines: numlines}

	if strings.HasPrefix(lines[0], "author") {
		authorline := splice(lines)
		bd.author = strings.TrimPrefix(authorline, "author ")
		authormail := splice(lines)
		bd.mail = strings.TrimPrefix(authormail, "author-mail ")
		authortime := splice(lines)
		bd.time = strings.TrimPrefix(authortime, "author-time")
		authortz := splice(lines)
		bd.time = strings.TrimPrefix(authortz, "author-tz")
	}

	return bd
}

func splice(sli []string) string {
	result := sli[0]
	sli = append(sli[:0], sli[1:]...)

	return result
}
