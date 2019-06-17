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
	oid        string
	Author     string
	NumLines   int
	Mail       string
	time       string
	tz         string
	otherLines []string
}

type commits map[string]BlameData

var Commits commits = make(commits)

func Parse(lines []string) []BlameData {

	inx := 0

	log.Debugf("beginning parse on lines: %v", lines)

	blameLines := BlameLines{lines: lines, indexPtr: inx}

	chunks := []BlameData{}

	for {

		blameLines.indexPtr = inx

		header, blameLines := ParseHeader(blameLines)
		extractedLines, blameLines := ParseLines(blameLines, header.NumLines)
		header.otherLines = append(header.otherLines, extractedLines...)

		chunks = append(chunks, header)

		inx = blameLines.indexPtr

		if blameLines.atEnd() {
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
		bd.Mail = strings.Trim(strings.Trim(strings.TrimPrefix(blines.shift(), "author-mail "), "<"), ">")
		bd.time = strings.TrimPrefix(blines.shift(), "author-time ")
		bd.tz = strings.TrimPrefix(blines.shift(), "author-tz ")
		Commits[bd.oid] = bd

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

	processLines := num*2 - 1

	log.Infof("processing %d lines starting with %s", processLines, lines.lines[lines.indexPtr])

	for i := 0; i < processLines; i++ {
		extracted = append(extracted, lines.shift())
	}

	return extracted, lines
}
