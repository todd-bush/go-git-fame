package parser

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"
)

type BlameLines struct {
	lines     []string
	index_ptr int
}

func shift(bl *BlameLines) string {
	result := current_line(*bl)

	if bl.index_ptr < len(bl.lines)-1 {
		bl.index_ptr = bl.index_ptr + 1
	}

	return result
}

func current_line(bl BlameLines) string {
	return bl.lines[bl.index_ptr]
}

func at_end(bl BlameLines) bool {
	return len(bl.lines)-1 == bl.index_ptr
}

type BlameData struct {
	oid         string
	author      string
	num_lines   int
	mail        string
	time        string
	tz          string
	other_lines []string
}

type commits map[string]BlameData

var Commits commits = make(commits)

func Parse(all_lines string) {

	lines := strings.Split(all_lines, "\n")

	inx := 0

	blame_lines := BlameLines{lines: lines, index_ptr: inx}

	chunks := []BlameData{}

	for {

		blame_lines.index_ptr = inx

		header, blame_lines := ParseHeader(blame_lines)
		extracted_lines, blame_lines := ParseLines(blame_lines, header.num_lines)
		header.other_lines = append(header.other_lines, extracted_lines...)

		chunks = append(chunks, header)

		inx = blame_lines.index_ptr

		if at_end(blame_lines) {
			break
		}

	}

	log.Debug(fmt.Sprintf("%v\n", chunks))

}

func ParseHeader(blame_lines BlameLines) (BlameData, BlameLines) {
	r, _ := regexp.Compile(`(?m)^([0-9a-f]{40}) (\d+) (\d+) (\d+)$`)

	headerline := shift(&blame_lines)

	pieces := r.FindStringSubmatch(headerline)
	numlines, _ := strconv.Atoi(pieces[4])

	bd := BlameData{oid: pieces[1], num_lines: numlines}

	if strings.HasPrefix(current_line(blame_lines), "author") {
		bd.author = strings.TrimPrefix(shift(&blame_lines), "author ")
		bd.mail = strings.TrimPrefix(shift(&blame_lines), "author-mail ")
		bd.time = strings.TrimPrefix(shift(&blame_lines), "author-time ")
		bd.tz = strings.TrimPrefix(shift(&blame_lines), "author-tz ")
		Commits[bd.oid] = bd
	} else {
		bd = Commits[bd.oid]
		bd.num_lines = numlines
	}

	// get to filename

	for {
		trash := shift(&blame_lines)
		if strings.HasPrefix(trash, "filename") || at_end(blame_lines) {
			break
		}
	}

	return bd, blame_lines
}

func ParseLines(lines BlameLines, num int) ([]string, BlameLines) {
	extracted := []string{}

	for i := 0; i < num; i++ {
		extracted = append(extracted, shift(&lines))
	}

	return extracted, lines
}
