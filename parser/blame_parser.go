package parser

import (
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type BlameLines struct {
	lines    []string
	indexPtr int
}

func (bl *BlameLines) shift() string {
	result := bl.currentLine()

	if bl.indexPtr < len(bl.lines)-1 {
		bl.indexPtr++
	}

	return result
}

func (bl BlameLines) currentLine() string {
	return bl.lines[bl.indexPtr]
}

func (bl BlameLines) atEnd() bool {
	return len(bl.lines)-1 == bl.indexPtr
}

type BlameData struct {
	oid         string
	Author      string
	NumLines    int
	Mail        string
	time        string
	tz          string
	other_lines []string
}

type commits map[string]BlameData

var Commits commits = make(commits)

func Parse(lines []string) []BlameData {

	inx := 0

	log.Debugf("beginning parse on lines: %v", lines)

	blame_lines := BlameLines{lines: lines, indexPtr: inx}

	chunks := []BlameData{}

	for {

		blame_lines.indexPtr = inx

		header, blame_lines := ParseHeader(blame_lines)
		extracted_lines, blame_lines := ParseLines(blame_lines, header.NumLines)
		header.other_lines = append(header.other_lines, extracted_lines...)

		chunks = append(chunks, header)

		inx = blame_lines.indexPtr

		if blame_lines.atEnd() {
			break
		}

	}

	return chunks
}

func ParseHeader(blines BlameLines) (BlameData, BlameLines) {
	r, _ := regexp.Compile(`(?m)^([0-9a-f]{40}) (\d+) (\d+) (\d+)$`)

	headerline := blines.shift()

	log.Infof("parsing header line: %s", headerline)

	pieces := r.FindStringSubmatch(headerline)
	numlines, _ := strconv.Atoi(pieces[4])

	bd := BlameData{oid: pieces[1], NumLines: numlines}

	if strings.HasPrefix(blines.currentLine(), "author") {
		bd.Author = strings.TrimPrefix(blines.shift(), "author ")
		bd.Mail = strings.TrimPrefix(blines.shift(), "author-mail ")
		bd.time = strings.TrimPrefix(blines.shift(), "author-time ")
		bd.tz = strings.TrimPrefix(blines.shift(), "author-tz ")
		Commits[bd.oid] = bd

		// clean up email
		bd.Mail = strings.Trim(strings.Trim(bd.Mail, "<"), ">")

		// get to filename

		for {
			trash := blines.shift()
			if strings.HasPrefix(trash, "filename") || blines.atEnd() {
				break
			}
		}

	} else {
		log.Debugf("using found data: %s", bd.oid)
		bd = Commits[bd.oid]
		log.Debugf("found bd data with author: %s", bd.Author)
		bd.NumLines = numlines
	}

	return bd, blines
}

func ParseLines(lines BlameLines, num int) ([]string, BlameLines) {
	extracted := []string{}

	process_lines := num*2 - 1

	log.Infof("processing %d lines starting with %s", process_lines, lines.lines[lines.indexPtr])

	for i := 0; i < process_lines; i++ {
		extracted = append(extracted, lines.shift())
	}

	return extracted, lines
}
