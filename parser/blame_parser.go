package parser

import (
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

var headerRegex = regexp.MustCompile(`(?m)^([0-9a-f]{40}) (\d+) (\d+) (\d+)$`)

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

func (bl *BlameLines) skip(num int) {
	bl.indexPtr += num
}

func (bl BlameLines) currentLine() string {
	return bl.lines[bl.indexPtr]
}

func (bl BlameLines) atEnd() bool {
	return len(bl.lines)-1 <= bl.indexPtr
}

type BlameData struct {
	oid      string
	Author   string
	NumLines int
	Mail     string
	time     string
	tz       string
}

type commits map[string]BlameData

var Commits commits = make(commits)

func Parse(lines []string) []BlameData {

	inx := 0

	log.Infof("beginning parse on lines: %v", lines)

	blameLines := BlameLines{lines: lines, indexPtr: inx}

	chunks := []BlameData{}

	for {

		blameLines.indexPtr = inx

		header, blameLines := ParseHeader(blameLines)
		blameLines = ParseLines(blameLines, header.NumLines)

		chunks = append(chunks, header)

		inx = blameLines.indexPtr

		if blameLines.atEnd() {
			break
		}

	}

	return chunks
}

func ParseHeader(blines BlameLines) (BlameData, BlameLines) {

	headerline := blines.shift()

	log.Infof("parsing header line: %s", headerline)

	pieces := headerRegex.FindStringSubmatch(headerline)
	numlines := 0

	if len(pieces) == 0 {
		log.Errorf("headerline doesn't match regex: %s", headerline)
	}

	if len(pieces) > 4 {
		numlines, _ = strconv.Atoi(pieces[4])
	}

	bd := BlameData{oid: pieces[1], NumLines: numlines}

	if strings.HasPrefix(blines.currentLine(), "author") {
		bd.Author = strings.TrimPrefix(blines.shift(), "author ")
		bd.Mail = strings.Trim(strings.Trim(strings.TrimPrefix(blines.shift(), "author-mail "), "<"), ">")
		bd.time = strings.TrimPrefix(blines.shift(), "author-time ")
		bd.tz = strings.TrimPrefix(blines.shift(), "author-tz ")
		Commits[bd.oid] = bd

		// move pointer to the end of this stanza
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

func ParseLines(lines BlameLines, num int) BlameLines {

	processLines := num*2 - 1

	// move the pointer this many
	lines.skip(processLines)

	return lines
}
