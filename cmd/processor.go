package cmd

import (
	log "github.com/sirupsen/logrus"
	git "github.com/todd-bush/go-git-fame/git_client"
	"github.com/todd-bush/go-git-fame/parser"
	"regexp"
	"strconv"
)

type BlameOutput struct {
	file       string
	blame_data []parser.BlameData
}

type ProcessOutput struct {
	author       string
	loc          int32
	commits      int32
	files        int32
	loc_perc     float32
	commits_perc float32
	files_perc   float32
}

type BlameProcess struct {
	file        string
	blame_lines []string
}

func ExecuteProcessor() []ProcessOutput {

	log.SetLevel(log.DebugLevel)

	result := []ProcessOutput{}

	// get this list of files
	file_list := git.GitListFiles("master") // TODO need to pass in branch

	log.Infof("found %d files to procesn", len(file_list))

	blame_out := []BlameProcess{}

	for _, file := range file_list {

		if len(file) > 0 {

			blame_result := git.GitBlame(file)
			out := BlameProcess{
				file:        file,
				blame_lines: blame_result,
			}

			blame_out = append(blame_out, out)
		}
	}

	blame_collector := []BlameOutput{}

	for _, bi := range blame_out {

		log.Infof("parsing blame on file: %s", bi.file)

		blame_out := parser.Parse(bi.blame_lines)
		blame_collector = append(blame_collector, BlameOutput{
			file:       bi.file,
			blame_data: blame_out,
		})
	}

	return result

}

func GatherCommits() map[string]int {
	result := map[string]int{}

	commit_lines := git.GitShortLog()

	r, _ := regexp.Compile(`(\d+)\s+(.+)\s+<(.+?)>`)

	for _, commit_line := range commit_lines {

		if len(commit_line) > 0 {

			log.Debugf("parsing line %s", commit_line)
			peices := r.FindStringSubmatch(commit_line)

			log.Debugf("peices = %v", peices)
			commits, _ := strconv.Atoi(peices[0])
			result[peices[2]] = commits
		}
	}

	return result
}
